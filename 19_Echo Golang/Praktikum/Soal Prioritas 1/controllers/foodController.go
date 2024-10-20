package controllers

import (
	"food-api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetFoods(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Foods)
}

func GetFoodByID(c echo.Context) error {
	id := c.Param("id")
	for _, food := range models.Foods {
		if food.ID == id {
			return c.JSON(http.StatusOK, food)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Makanan tidak ditemukan"})
}

func CreateFood(c echo.Context) error {
	newFood := new(models.Food)
	if err := c.Bind(newFood); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Permintaan tidak valid"})
	}

	newFood.ID = uuid.New().String()
	models.Foods = append(models.Foods, *newFood)
	return c.JSON(http.StatusCreated, newFood)
}

func UpdateFood(c echo.Context) error {
	id := c.Param("id")
	updatedFood := new(models.Food)
	if err := c.Bind(updatedFood); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Permintaan tidak valid"})
	}

	for i, food := range models.Foods {
		if food.ID == id {
			updatedFood.ID = id
			models.Foods[i] = *updatedFood
			return c.JSON(http.StatusOK, updatedFood)
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{"message": "Makanan tidak ditemukan"})
}

func DeleteFood(c echo.Context) error {
	id := c.Param("id")
	for i, food := range models.Foods {
		if food.ID == id {
			models.Foods = append(models.Foods[:i], models.Foods[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{"message": "Makanan berhasil dihapus"})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Makanan tidak ditemukan"})
}
