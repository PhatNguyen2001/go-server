package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString(os.Getenv("SECRET_KEY"))

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
