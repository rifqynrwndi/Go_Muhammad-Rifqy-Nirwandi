package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"app/models"
)

type PackageController struct {
	DB *gorm.DB
}

func (pc *PackageController) GetAllPackages(c echo.Context) error {
	var packages []models.Package
	pc.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func (pc *PackageController) GetPackageByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pack models.Package
	if err := pc.DB.First(&pack, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "package not found"})
	}
	return c.JSON(http.StatusOK, pack)
}

func (pc *PackageController) CreatePackage(c echo.Context) error {
	pack := new(models.Package)
	if err := c.Bind(pack); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Create(pack)
	return c.JSON(http.StatusCreated, pack)
}

func (pc *PackageController) UpdatePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pack models.Package
	if err := pc.DB.First(&pack, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "package not found"})
	}
	if err := c.Bind(&pack); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pc.DB.Save(&pack)
	return c.JSON(http.StatusOK, pack)
}

func (pc *PackageController) DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.DB.Delete(&models.Package{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "package not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
