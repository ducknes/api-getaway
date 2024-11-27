package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

func GetUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := client.GetUser(c.Context(), c.Query("id"))
		if err != nil {
			return err
		}

		return c.JSON(user)
	}
}
