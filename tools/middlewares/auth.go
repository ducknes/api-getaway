package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
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

	fmt.Println(os.Getenv("JWT_SECRET"))

	jwtToken, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	tokenClaims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return errors.New("invalid token")
	}

	exp, ok := tokenClaims["exp"].(float64)
	if !ok {
		return errors.New("invalid token")
	}

	if time.Now().Unix() > int64(exp) {
		return errors.New("token expired")
	}

	return nil
}
