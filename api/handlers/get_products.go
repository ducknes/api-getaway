package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetProductsHandler получает продукты.
//
// @Summary Получение продуктов
// @Description Получает определенное через limit количество продуктов
// @Tags Продукты
// @Produce json
// @Param limit query int true "Количество получаемых продуктов"
// @Param cursor query string false "Ссылка на следующие продукты"
// @Success 200 {object} storageservice.Products "Получение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products [get]
// @Security LogisticAuth
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
