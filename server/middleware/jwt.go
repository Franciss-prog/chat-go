package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UUID     string
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(uuid, username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := Claims{
		UUID:     uuid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func AuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}
	claims, err := ValidateToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}
	c.Locals("userID", claims.UUID)
	c.Locals("username", claims.Username)
	return c.Next()
}
