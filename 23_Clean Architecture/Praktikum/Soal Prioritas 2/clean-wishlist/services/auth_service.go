package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
	secretKey string
}

func NewAuthService(secretKey string) AuthService {
	return &authService{secretKey: secretKey}
}

func (s *authService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
