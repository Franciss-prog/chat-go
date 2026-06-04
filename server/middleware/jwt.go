package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UUID string
	jwt.RegisteredClaims
}

func GenerateToken(uuid string) (string, error) {
	// get the token from env
	secret := os.Getenv("JWT_SECRET")

	// create the claims using Claims struct
	claims := Claims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// register the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return  the signed token
	return token.SignedString([]byte(secret))
}
