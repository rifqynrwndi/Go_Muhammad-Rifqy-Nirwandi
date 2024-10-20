package main

import (
	"word-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.WordRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
