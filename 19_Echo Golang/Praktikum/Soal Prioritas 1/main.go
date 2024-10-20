package main

import (
	"food-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.FoodRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
