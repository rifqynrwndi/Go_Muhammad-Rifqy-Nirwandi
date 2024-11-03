package controllers

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	service services.WishlistService
}

func NewWishlistController(service services.WishlistService) *WishlistController {
	return &WishlistController{service: service}
}

func (ctrl *WishlistController) GetAll(c echo.Context) error {
	wishlists, err := ctrl.service.GetAllWishlists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": wishlists})
}

func (ctrl *WishlistController) Create(c echo.Context) error {
	var wishlist entities.Wishlist
	if err := c.Bind(&wishlist); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := ctrl.service.CreateWishlist(&wishlist); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": wishlist})
}
