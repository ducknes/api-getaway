package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

// SaveProductsHandler сохраняет продукты.
//
// @Summary Сохранение продуктов
// @Description Сохраняет продукты
// @Tags Продукты
// @Produce json
// @Param products body []storageservice.Product true "Сохраняемые продукты"
// @Success 200 "Сохранение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products [post]
// @Security LogisticAuth
func SaveProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var products []storageservice.Product
		if err := c.BodyParser(&products); err != nil {
			return err
		}

		return client.SaveProducts(c.Context(), products)
	}
}
