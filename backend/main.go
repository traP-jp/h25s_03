package main

import (
	"os"

	"github.com/eraxyso/go-template/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	OapiValidator "github.com/oapi-codegen/echo-middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	swagger, err := api.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(OapiValidator.OapiRequestValidator(swagger))

	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		e.Logger.Fatal("DB_USER environment variable is not set")
	}
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		e.Logger.Fatal("DB_PASSWORD environment variable is not set")
	}
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		e.Logger.Fatal("DB_HOST environment variable is not set")
	}
	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		e.Logger.Fatal("DB_PORT environment variable is not set")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		e.Logger.Fatal("DB_NAME environment variable is not set")
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	server := InitializeServer(db)

	mr := newMiddlewareRouter()

	mr.addRoute("POST", "/exmaples", server.MiddlewareService.MiddlewareServiceExample)
	mr.addGroup("POST", "/examples", server.MiddlewareService.MiddlewareServiceExample)

	e.Use(mr.registerMiddlewares)

	api.RegisterHandlers(e, server)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		e.Logger.Fatal("PORT environment variable is not set")
	}
	e.Logger.Fatal(e.Start(port))
}
