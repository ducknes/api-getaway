package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, cursor := parseQueryParams(c)

		products, err := client.GetProducts(c.Context(), limit, cursor)
		if err != nil {
			return err
		}

		return c.JSON(products)
	}
}

func parseQueryParams(ctx *fiber.Ctx) (int, string) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 0
	}

	return limit, ctx.Query("cursor")
}
