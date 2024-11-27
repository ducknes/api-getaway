package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

func GetProductHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product, err := client.GetProduct(c.Context(), c.Query("id"))
		if err != nil {
			return err
		}

		return c.JSON(&product)
	}
}
