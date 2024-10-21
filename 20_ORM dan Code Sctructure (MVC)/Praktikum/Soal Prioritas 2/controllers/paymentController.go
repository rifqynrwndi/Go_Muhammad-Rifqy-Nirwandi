package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"billing/models"
)

type PaymentController struct {
	DB *gorm.DB
}

func (pc *PaymentController) GetAllPayments(c echo.Context) error {
	var payments []models.Payment
	pc.DB.Find(&payments)
	return c.JSON(http.StatusOK, payments)
}

func (pc *PaymentController) GetPaymentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var payment models.Payment
	if err := pc.DB.First(&payment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "payment not found"})
	}
	return c.JSON(http.StatusOK, payment)
}

func (pc *PaymentController) CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Create(payment)
	return c.JSON(http.StatusCreated, payment)
}

func (pc *PaymentController) UpdatePayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var payment models.Payment
	if err := pc.DB.First(&payment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "payment not found"})
	}
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Save(&payment)
	return c.JSON(http.StatusOK, payment)
}

func (pc *PaymentController) DeletePayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.DB.Delete(&models.Payment{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "payment not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
