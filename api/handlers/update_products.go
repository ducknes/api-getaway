package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

// UpdateProductsHandler обновляет продукты.
//
// @Summary Обновление продуктов
// @Description Обновляет продукты
// @Tags Продукты
// @Param products body []storageservice.Product true "Обновляемые продукты"
// @Success 200 "Сохранение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products [put]
// @Security LogisticAuth
func UpdateProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var products []storageservice.Product
		if err := c.BodyParser(&products); err != nil {
			return err
		}

		return client.UpdateProducts(c.Context(), products)
	}
}
