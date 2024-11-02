package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"app/controllers"
)

func SetupRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	packageController := controllers.PackageController{DB: db}

	e.GET("/api/v1/packages", packageController.GetAllPackages)
	e.GET("/api/v1/packages/:id", packageController.GetPackageByID)
	e.POST("/api/v1/packages", packageController.CreatePackage)
	e.PUT("/api/v1/packages/:id", packageController.UpdatePackage)
	e.DELETE("/api/v1/packages/:id", packageController.DeletePackage)

	return e
}