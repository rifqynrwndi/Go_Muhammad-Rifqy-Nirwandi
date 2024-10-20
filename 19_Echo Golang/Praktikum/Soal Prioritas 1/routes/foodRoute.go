package routes

import (
	"food-api/controllers"
	"github.com/labstack/echo/v4"
)

func FoodRoutes(e *echo.Echo) {
	e.GET("/api/v1/foods", controllers.GetFoods)
	e.GET("/api/v1/foods/:id", controllers.GetFoodByID)
	e.POST("/api/v1/foods", controllers.CreateFood)
	e.PUT("/api/v1/foods/:id", controllers.UpdateFood)
	e.DELETE("/api/v1/foods/:id", controllers.DeleteFood)
}
