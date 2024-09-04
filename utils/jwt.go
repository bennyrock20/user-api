package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var stringKey = GetEnv("stringKey", "")

var jwtSecret = []byte(stringKey)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
