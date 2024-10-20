package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddBookHandler(c echo.Context) error {
	book := models.Books{}
	book.Title =c.FormValue("title")
	c.Bind(&book)
	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,  models.BaseResponse{false, "Failed to Connect", nil})
	}
	return c.JSON(200, models.BaseResponse{true, "Sukses", book})
}

func GetDetailBooksHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	book := models.Books{
		idNumber, "Buku Fisika", "Intan",
	}
	return c.JSON(200, models.BaseResponse{true, "Sukses", book})
}

func GetBooksHandler(c echo.Context) error {
	books := []models.Books{}
	result := config.DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,  models.BaseResponse{false, "Failed to Get Data", nil})
	}
	return c.JSON(200, models.BaseResponse{true, "Sukses", books})
}