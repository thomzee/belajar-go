package main

import (
	"assignment_2/configs"
	"assignment_2/docs"
	"assignment_2/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"os"
	"strconv"
)

func main() {
	// App environment start here ...
	configs.Env()
	appPort := os.Getenv("APP_PORT")

	// Database initialization
	configs.InitDb()
	db := configs.GetDb()

	// Redis initialization
	configs.InitRedis()

	// Schema migration start here ...
	reset, _ := strconv.ParseBool(os.Getenv("DB_RESET"))
	configs.Migrate(db, reset)

	// Echo initialization
	e := echo.New()

	// App global middleware start here ...
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// App routes start here ...
	routes.ApiRoute(e, db)

	// Swagger start here ...
	docs.SwaggerInfo.Title = "Assignment 2 Docs"
	docs.SwaggerInfo.Description = "This is API documentation of assignment 2nd"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:"+appPort
	docs.SwaggerInfo.Schemes = []string{"http"}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// App runner start here ...
	e.Logger.Fatal(e.Start(":" + appPort))
}
