package controllers

import (
    "net/http"
    "strconv"
    "project/config"
    "project/models"
    "github.com/labstack/echo/v4"
)

func GetPackagesHandler(c echo.Context) error {
    var packages []models.Package
    if err := config.DB.Find(&packages).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not fetch packages"})
    }
    return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Data: packages})
}

func GetPackageByIDHandler(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var pkg models.Package
    if err := config.DB.First(&pkg, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, models.BaseResponse{Status: "fail", Message: "Package not found"})
    }
    return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Data: pkg})
}

func AddPackageHandler(c echo.Context) error {
    var pkg models.Package
    if err := c.Bind(&pkg); err != nil {
        return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input"})
    }
    if err := config.DB.Create(&pkg).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not create package"})
    }
    return c.JSON(http.StatusCreated, models.BaseResponse{Status: "success", Message: "Package added successfully"})
}

func UpdatePackageHandler(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var pkg models.Package
    if err := config.DB.First(&pkg, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, models.BaseResponse{Status: "fail", Message: "Package not found"})
    }
    if err := c.Bind(&pkg); err != nil {
        return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input"})
    }
    if err := config.DB.Save(&pkg).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not update package"})
    }
    return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Package updated successfully"})
}

func DeletePackageHandler(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := config.DB.Delete(&models.Package{}, id).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not delete package"})
    }
    return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Package deleted successfully"})
}
