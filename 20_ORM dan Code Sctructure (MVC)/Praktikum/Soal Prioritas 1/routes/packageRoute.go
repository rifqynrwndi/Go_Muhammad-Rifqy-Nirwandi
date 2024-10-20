package routes

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "app/controllers"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
    controllers.InitializeDB(db)

    e.GET("/api/v1/packages", controllers.GetAllPackages)
    e.GET("/api/v1/packages/:id", controllers.GetPackageByID)
    e.POST("/api/v1/packages", controllers.CreatePackage)
    e.PUT("/api/v1/packages/:id", controllers.UpdatePackage)
    e.DELETE("/api/v1/packages/:id", controllers.DeletePackage)
}
