package routes

import (
	"word-api/controllers"
	"github.com/labstack/echo/v4"
)

func WordRoutes(e *echo.Echo) {
	e.GET("/words", controllers.GetWords)
	e.POST("/words", controllers.AddWord)
}
