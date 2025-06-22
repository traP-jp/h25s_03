package main

import (
	"net/http"
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
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	server := InitializeServer(db)

	mr := newMiddlewareRouter()

	mr.addRoute(http.MethodPatch, "/events/:eventID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodPost, "/events/:eventID/lotteries", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodPost, "/events/:eventID/:lotteryID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID/:lotteryID", server.MiddlewareService.EventAdminAuthentication)

	mr.addRoute(http.MethodPost, "/events/:eventID/attendance", server.MiddlewareService.EventRegistrationAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID/attendance", server.MiddlewareService.EventRegistrationAuthentication)

	e.Use(mr.registerMiddlewares)

	api.RegisterHandlers(e, server)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		e.Logger.Fatal("PORT environment variable is not set")
	}
	e.Logger.Fatal(e.Start(port))
}
