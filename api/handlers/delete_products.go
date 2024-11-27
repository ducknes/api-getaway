package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

func DeleteProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var productIds []string
		if err := ctx.BodyParser(&productIds); err != nil {
			return err
		}

		return client.DeleteProducts(ctx.Context(), productIds)
	}
}
