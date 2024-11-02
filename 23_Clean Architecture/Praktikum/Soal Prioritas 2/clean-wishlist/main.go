package main

import (
	"go-wishlist-api/config"
	"go-wishlist-api/controllers"
	"go-wishlist-api/middleware"
	"go-wishlist-api/repositories"
	"go-wishlist-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDatabase()
	db := config.DB

	wishlistRepo := repositories.NewWishlistRepository(db)
	wishlistService := services.NewWishlistService(wishlistRepo)
	wishlistController := controllers.NewWishlistController(wishlistService)

	e := echo.New()

	// Routes
	e.POST("/login", func(c echo.Context) error {
		token, err := middleware.GenerateToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"token": token})
	})

	protectedRoutes := e.Group("/wishlists")
	protectedRoutes.Use(middleware.JWTMiddleware)
	protectedRoutes.GET("", wishlistController.GetAll)
	protectedRoutes.POST("", wishlistController.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
