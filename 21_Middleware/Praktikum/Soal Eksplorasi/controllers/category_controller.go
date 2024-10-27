package controllers

import (
	"net/http"
	"strconv"
	"postcategory/config"
	"postcategory/models"
	"github.com/labstack/echo/v4"
)

func GetCategoriesHandler(c echo.Context) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not fetch categories"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Data: categories})
}

func AddCategoryHandler(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input"})
	}
	if err := config.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not create category"})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{Status: "success", Message: "Category added successfully"})
}

func DeleteCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not delete category"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Category deleted successfully"})
}
