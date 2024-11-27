package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

func SaveProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var products []storageservice.Product
		if err := c.BodyParser(&products); err != nil {
			return err
		}

		return client.SaveProducts(c.Context(), products)
	}
}
