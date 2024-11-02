package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("rifqy")

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		return next(c)
	}
}

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
