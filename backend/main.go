package main

import (
	"net/http"
	"os"
	"time"

	"traquji/api"
	"traquji/repository"

	sql_mysql "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// not working correctly on noeshowcase, to be fixed
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

	config := &sql_mysql.Config{
		User:                 dbUser,
		Passwd:               dbPassword,
		Net:                  "tcp",
		Addr:                 dbHost + ":" + dbPort,
		DBName:               dbName,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	db, err := gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database:", err)
	}
	err = repository.Migrate(db)
	if err != nil {
		e.Logger.Fatal("Failed to migrate database:", err)
	}
	server := InitializeServer(db)

	mr := newMiddlewareRouter()

	mr.addRoute(http.MethodPatch, "/events/:eventID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodPost, "/events/:eventID/lotteries", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodPost, "/events/:eventID/lotteries/:lotteryID", server.MiddlewareService.EventAdminAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID/lotteries/:lotteryID", server.MiddlewareService.EventAdminAuthentication)

	mr.addRoute(http.MethodPost, "/events/:eventID/attendance", server.MiddlewareService.EventRegistrationAuthentication)
	mr.addRoute(http.MethodDelete, "/events/:eventID/attendance", server.MiddlewareService.EventRegistrationAuthentication)

	e.Use(mr.registerMiddlewares)

	api.RegisterHandlers(e, server)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		e.Logger.Fatal("PORT environment variable is not set")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
