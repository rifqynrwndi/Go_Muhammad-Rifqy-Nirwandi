package main

import (
	"rest/config"
	"rest/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDB()
	e := echo.New()
	e.GET("/books", controllers.GetBooksHandler)
	e.GET("/books/:id", controllers.GetDetailBooksHandler)
	e.POST("/books", controllers.AddBookHandler)
	e.Start(":8000")
}


