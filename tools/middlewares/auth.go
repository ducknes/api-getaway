package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if len(token) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	if err := validateToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	return c.Next()
}

func validateToken(token string) error {
	splitToken := strings.Split(token, " ")

	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return errors.New("invalid token")
	}

	_, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		return err
	}

	return nil
}
