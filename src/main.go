package main

import (
	"base/config"
	"base/databases"
	"base/routers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading '.env' file")
	}
}

func main() {
	// middleware
	m := middlewares()

	// config
	c := config.Config{}

	// database
	d := databases.ConnectMySQL()
	d.AutoMigrate()

	// router
	v1 := m.Group("/api")
	routers.Lists(v1)

	m.Logger.Fatal(m.Start(c.AppConfig().URL + ":" + c.AppConfig().Port))
}

func middlewares() *echo.Echo {
	e := echo.New()
	// e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
