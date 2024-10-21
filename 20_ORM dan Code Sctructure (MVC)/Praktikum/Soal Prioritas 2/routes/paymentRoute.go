package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"billing/controllers"
)

func SetupRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	paymentController := controllers.PaymentController{DB: db}

	e.GET("/api/v1/payments", paymentController.GetAllPayments)
	e.GET("/api/v1/payments/:id", paymentController.GetPaymentByID)
	e.POST("/api/v1/payments", paymentController.CreatePayment)
	e.PUT("/api/v1/payments/:id", paymentController.UpdatePayment)
	e.DELETE("/api/v1/payments/:id", paymentController.DeletePayment)

	return e
}
