package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

func SaveUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userservice.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		return client.SaveUser(c.Context(), user)
	}
}
