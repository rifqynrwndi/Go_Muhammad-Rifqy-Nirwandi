package controllers

import (
	"log"
	"net/http"
	"project/config"
	"project/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input", Data: nil})
	}

	oldPassword := user.Password
	result := config.DB.First(&user, "email = ?", user.Email)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: "fail", Message: "Email not found", Data: nil})
	}

	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: "fail", Message: "Incorrect password", Data: nil})
	}

	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Failed to generate token", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "Login successful", Data: token})
}

func RegisterHandler(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: "fail", Message: "Invalid input", Data: nil})
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Failed to hash password", Data: nil})
	}
	user.Password = hash

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: "fail", Message: "Failed to register user", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: "success", Message: "User registered successfully", Data: user})
}

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Request: %s %s", c.Request().Method, c.Request().URL)
		return next(c)
	}
}

func GenerateJWT(userID uint, email string) (string, error) {
	claims := &jwtCustomClaims{
		Email:  email,
		UserID: int(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("rifqy"))
	if err != nil {
		return "", err
	}

	return t, nil
}
