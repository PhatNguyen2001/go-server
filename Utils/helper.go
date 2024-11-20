package helper

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
)

func GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("SECRET_KEY")
	refreshToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func GenerateTemporaryPassword(length int) string {
	const characters = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func SendResetEmail(email string, tempPassword string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))                                              // Replace with your email
	m.SetHeader("To", email)                                                             // Email recipient
	m.SetHeader("Subject", "Password Reset Request")                                     // Email subject
	m.SetBody("text/plain", fmt.Sprintf("Your temporary password is: %s", tempPassword)) // Email body

	// SMTP Configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	// Send email
	return d.DialAndSend(m)
}
