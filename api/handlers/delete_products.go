package handlers

import (
	"api-getaway/cluster/storageservice"
	"github.com/gofiber/fiber/v2"
)

// DeleteProductsHandler удаляет продукты по списку их идентификаторов.
//
// @Summary Удаление продуктов
// @Description Удаляет продукты из базы данных по переданным идентификаторам
// @Tags Продукты
// @Param productIds body []string true "Список идентификаторов продуктов для удаления"
// @Success 200 "Удаление прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /products [delete]
// @Security LogisticAuth
func DeleteProductsHandler(client *storageservice.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var productIds []string
		if err := ctx.BodyParser(&productIds); err != nil {
			return err
		}

		return client.DeleteProducts(ctx.Context(), productIds)
	}
}
