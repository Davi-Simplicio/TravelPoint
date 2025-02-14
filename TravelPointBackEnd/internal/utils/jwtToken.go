package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("$Ecr3t")

func GenerateJWTToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":         userID,
		"exp":             time.Now().Add(24 * time.Hour).Unix(),
		"isAuthenticated": false,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
