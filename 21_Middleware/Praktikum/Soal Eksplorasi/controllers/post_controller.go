package controllers

import (
	"net/http"
	"strconv"
	"postcategory/config"
	"postcategory/models"
	"github.com/labstack/echo/v4"
)

func GetPostsHandler(c echo.Context) error {
	var posts []models.Post
	if err := config.DB.Preload("Category").Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not fetch posts"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Data: posts})
}

func GetPostByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := config.DB.Preload("Category").First(&post, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: "fail", Message: "Post not found"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Data: post})
}

func AddPostHandler(c echo.Context) error {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input"})
	}
	if err := config.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not create post"})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{Status: "success", Message: "Post added successfully"})
}

func UpdatePostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: "fail", Message: "Post not found"})
	}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input"})
	}
	post.ID = uint(id)
	if err := config.DB.Save(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not update post"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Post updated successfully"})
}

func DeletePostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Could not delete post"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Post deleted successfully"})
}
