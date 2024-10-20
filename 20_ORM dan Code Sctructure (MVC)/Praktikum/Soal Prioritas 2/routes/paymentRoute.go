package routes

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "billing/controllers"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
    controllers.InitializeDB(db)

    e.GET("/api/v1/payments", controllers.GetAllPayments)
    e.GET("/api/v1/payments/:id", controllers.GetPaymentByID)
    e.POST("/api/v1/payments", controllers.CreatePayment)
    e.PUT("/api/v1/payments/:id", controllers.UpdatePayment)
    e.DELETE("/api/v1/payments/:id", controllers.DeletePayment)
}
