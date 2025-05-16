package user

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(login string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
