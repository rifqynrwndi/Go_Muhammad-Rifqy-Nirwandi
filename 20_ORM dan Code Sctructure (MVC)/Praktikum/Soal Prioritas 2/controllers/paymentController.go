package controllers

import (
    "database/sql"
    "net/http"
    "github.com/labstack/echo/v4"
    "billing/models"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
    db = database
}

func GetAllPayments(c echo.Context) error {
    rows, err := db.Query("SELECT * FROM payments")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to fetch payments"})
    }
    defer rows.Close()

    payments := []models.Payment{}
    for rows.Next() {
        var payment models.Payment
        if err := rows.Scan(&payment.ID, &payment.Title, &payment.Description, &payment.Amount); err != nil {
            return err
        }
        payments = append(payments, payment)
    }
    return c.JSON(http.StatusOK, payments)
}

func GetPaymentByID(c echo.Context) error {
    id := c.Param("id")
    var payment models.Payment
    err := db.QueryRow("SELECT * FROM payments WHERE id = ?", id).Scan(&payment.ID, &payment.Title, &payment.Description, &payment.Amount)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Payment not found"})
    }
    return c.JSON(http.StatusOK, payment)
}

func CreatePayment(c echo.Context) error {
    var payment models.Payment
    if err := c.Bind(&payment); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    result, err := db.Exec("INSERT INTO payments (title, description, amount) VALUES (?, ?, ?)", payment.Title, payment.Description, payment.Amount)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create payment"})
    }
    id, _ := result.LastInsertId()
    payment.ID = int(id)
    return c.JSON(http.StatusCreated, payment)
}

func UpdatePayment(c echo.Context) error {
    id := c.Param("id")
    var payment models.Payment
    if err := c.Bind(&payment); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    _, err := db.Exec("UPDATE payments SET title = ?, description = ?, amount = ? WHERE id = ?", payment.Title, payment.Description, payment.Amount, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update payment"})
    }
    return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) error {
    id := c.Param("id")
    _, err := db.Exec("DELETE FROM payments WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Payment not found"})
    }
    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully"})
}
