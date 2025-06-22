package main

import (
	"net/url"
	"os"

	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// FIXME: バリデーションが正常なものも弾いてしまうのでコメントアウト
	// swagger, err := api.GetSwagger()
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// e.Use(OapiValidator.OapiRequestValidator(swagger))

	dbUser, exists := os.LookupEnv("NS_MARIADB_USER")
	if !exists {
		e.Logger.Fatal("NS_MARIADB_USER environment variable is not set")
	}
	dbPassword, exists := os.LookupEnv("NS_MARIADB_PASSWORD")
	if !exists {
		e.Logger.Fatal("NS_MARIADB_PASSWORD environment variable is not set")
	}
	dbHost, exists := os.LookupEnv("NS_MARIADB_HOSTNAME")
	if !exists {
		e.Logger.Fatal("NS_MARIADB_HOSTNAME environment variable is not set")
	}
	dbPort, exists := os.LookupEnv("NS_MARIADB_PORT")
	if !exists {
		e.Logger.Fatal("NS_MARIADB_PORT environment variable is not set")
	}
	dbName, exists := os.LookupEnv("NS_MARIADB_DATABASE")
	if !exists {
		e.Logger.Fatal("NS_MARIADB_DATABASE environment variable is not set")
	}

	escapedPassword := url.QueryEscape(dbPassword)
	dsn := dbUser + ":" + escapedPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database:", err)
	}
	err = repository.Migrate(db)
	if err != nil {
		e.Logger.Fatal("Failed to migrate database:", err)
	}
	server := InitializeServer(db)

	mr := newMiddlewareRouter()

	// TODO: Add admin routes
	// mr.addRoute("POST", "/exmaples", server.MiddlewareService.MiddlewareServiceExample)
	// mr.addGroup("POST", "/examples", server.MiddlewareService.MiddlewareServiceExample)

	e.Use(mr.registerMiddlewares)

	api.RegisterHandlers(e, server)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		e.Logger.Fatal("PORT environment variable is not set")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
