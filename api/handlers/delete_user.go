package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

// DeleteUserHandler удаляет пользователя по его идентификатору.
//
// @Summary Удаление пользователя
// @Description Удаляет пользователя по его идентификатору
// @Tags Пользователи
// @Param id query string true "Идентификатор пользователя"
// @Success 200 "Удаление прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /user [delete]
// @Security LogisticAuth
func DeleteUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return client.DeleteUser(c.Context(), c.Query("id"))
	}
}
