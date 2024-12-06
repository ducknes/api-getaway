package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

// GetProductHandler получает продукт по его идентификатору.
//
// @Summary Получение продукта
// @Description Получает продукт по его идентификатору
// @Tags Продукты
// @Produce json
// @Param id query string true "Идентификатор продукта"
// @Success 200 {object} storageservice.Product "Получение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products/one [get]
// @Security LogisticAuth
func GetProductHandler(client *storageservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product, err := client.GetProduct(c.Context(), c.Query("id"))
		if err != nil {
			return err
		}

		return c.JSON(&product)
	}
}
