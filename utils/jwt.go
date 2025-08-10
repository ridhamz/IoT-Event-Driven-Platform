package utils

import (
	"go-cqrs-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.AppConfig.JWT_SECRET) // replace with your own secret

func GenerateJWT(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // token expires in 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
