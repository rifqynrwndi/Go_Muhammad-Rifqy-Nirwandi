package controllers

import (
	"crypto/sha256"
	"rest/config"
	"rest/models"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context,) error {
	user := models.User{}
	c.Bind(&user)
	h := sha256.New()
	h.Write([]byte(user.Password))
	bs := h.Sum(nil)

	result := config.DB.First(&user, "email = ? AND password = ?", user.Email, bs)
	if result.Error != nil {
		if user.ID == 0 {
			return c.JSON(400, map[string]interface{}{
				"message": "Email atau Password salah",
			})
		}

		return c.JSON(500, map[string]interface{}{
			"message": "Database Error",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Login Success",
		"user":    user,
	})

}
