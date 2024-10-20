package routes

import (
    "github.com/labstack/echo/v4"
    "word-api/controllers"
)

func SetupRoutes(e *echo.Echo) {
    e.GET("/api/v1/words", controllers.GetAllWords)
    e.GET("/api/v1/words/:id", controllers.GetWordByID)
    e.POST("/api/v1/words", controllers.AddWord)
    e.DELETE("/api/v1/words/:id", controllers.DeleteWord)
}
