package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, cursor := parseQueryParams(c)

		users, err := client.GetUsers(c.Context(), limit, cursor)
		if err != nil {
			return err
		}

		return c.JSON(users)
	}
}
