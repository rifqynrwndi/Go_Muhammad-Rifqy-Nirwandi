package controllers

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	Service services.WishlistService
}

func NewWishlistController(service services.WishlistService) *WishlistController {
	return &WishlistController{Service: service}
}

func (c *WishlistController) GetAll(ctx echo.Context) error {
	wishlists, err := c.Service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"data": wishlists})
}

func (c *WishlistController) Create(ctx echo.Context) error {
	var input entities.Wishlist
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	wishlist, err := c.Service.Create(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"data": wishlist})
}
