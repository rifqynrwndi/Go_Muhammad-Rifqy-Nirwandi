package controllers

import (
	"app/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
    db = database
}

func GetAllPackages(c echo.Context) error {
    rows, err := db.Query("SELECT * FROM packages")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to fetch packages"})
    }
    defer rows.Close()

    packages := []models.Package{}
    for rows.Next() {
        var pkg models.Package
        if err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.Sender, &pkg.Receiver, &pkg.SenderLocation, &pkg.ReceiverLocation, &pkg.Fee, &pkg.Weight); err != nil {
            return err
        }
        packages = append(packages, pkg)
    }
    return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
    id := c.Param("id")
    var pkg models.Package
    err := db.QueryRow("SELECT * FROM packages WHERE id = ?", id).Scan(&pkg.ID, &pkg.Name, &pkg.Sender, &pkg.Receiver, &pkg.SenderLocation, &pkg.ReceiverLocation, &pkg.Fee, &pkg.Weight)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
    }
    return c.JSON(http.StatusOK, pkg)
}

func CreatePackage(c echo.Context) error {
    var pkg models.Package
    if err := c.Bind(&pkg); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    result, err := db.Exec("INSERT INTO packages (name, sender, receiver, sender_location, receiver_location, fee, weight) VALUES (?, ?, ?, ?, ?, ?, ?)", pkg.Name, pkg.Sender, pkg.Receiver, pkg.SenderLocation, pkg.ReceiverLocation, pkg.Fee, pkg.Weight)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create package"})
    }
    id, _ := result.LastInsertId()
    pkg.ID = int(id)
    return c.JSON(http.StatusCreated, pkg)
}

func UpdatePackage(c echo.Context) error {
    id := c.Param("id")
    var pkg models.Package
    if err := c.Bind(&pkg); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    _, err := db.Exec("UPDATE packages SET name = ?, sender = ?, receiver = ?, sender_location = ?, receiver_location = ?, fee = ?, weight = ? WHERE id = ?", pkg.Name, pkg.Sender, pkg.Receiver, pkg.SenderLocation, pkg.ReceiverLocation, pkg.Fee, pkg.Weight, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update package"})
    }
    return c.JSON(http.StatusOK, pkg)
}

func DeletePackage(c echo.Context) error {
    id := c.Param("id")
    _, err := db.Exec("DELETE FROM packages WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
    }
    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully"})
}
