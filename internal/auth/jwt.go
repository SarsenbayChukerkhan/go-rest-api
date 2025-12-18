package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("SuperSecretKey123456")

func Generate(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
