package main

import (
	"posts/config"
	"posts/controllers"

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

	eAuth.GET("/api/v1/posts", controllers.GetPostsHandler)
	eAuth.GET("/api/v1/posts/:id", controllers.GetPostByIDHandler)
	eAuth.POST("/api/v1/posts", controllers.AddPostHandler)
	eAuth.PUT("/api/v1/posts/:id", controllers.UpdatePostHandler)
	eAuth.DELETE("/api/v1/posts/:id", controllers.DeletePostHandler)

	e.POST("/api/v1/login", controllers.LoginHandler)
	e.POST("/api/v1/register", controllers.RegisterHandler)

	e.Start(":8000")
}
