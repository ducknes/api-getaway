package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

func DeleteUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return client.DeleteUser(c.Context(), c.Query("id"))
	}
}
