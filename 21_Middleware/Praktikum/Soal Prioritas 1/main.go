package main

import (
	"project/config"
	"project/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDB()

	e := echo.New()
	e.Use(controllers.LogMiddleware)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("rifqy")))

	eAuth.GET("/api/v1/packages", controllers.GetPackagesHandler)
	eAuth.GET("/api/v1/packages/:id", controllers.GetPackageByIDHandler)
	eAuth.POST("/api/v1/packages", controllers.AddPackageHandler)
	eAuth.PUT("/api/v1/packages/:id", controllers.UpdatePackageHandler)
	eAuth.DELETE("/api/v1/packages/:id", controllers.DeletePackageHandler)

	e.POST("/api/v1/login", controllers.LoginHandler)
	e.POST("/api/v1/register", controllers.RegisterHandler)

	e.Start(":8000")
}
