package main

import (
	"go-wishlist-api/config"
	"go-wishlist-api/controllers"
	"go-wishlist-api/repositories"
	"go-wishlist-api/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDatabase()

	db := config.DB
	wishlistRepo := repositories.NewWishlistRepository(db)
	wishlistService := services.NewWishlistService(wishlistRepo)
	wishlistController := controllers.NewWishlistController(wishlistService)

	e := echo.New()
	e.GET("/wishlists", wishlistController.GetAll)
	e.POST("/wishlists", wishlistController.Create)

	e.Logger.Fatal(e.Start(":8000"))
}
