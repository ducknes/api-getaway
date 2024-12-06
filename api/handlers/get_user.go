package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

// GetUserHandler получает пользователя по его идентификатору.
//
// @Summary Получение пользователя
// @Description Получает пользователя по его идентификатору
// @Tags Пользователи
// @Produce json
// @Param id query string true "Идентификатор пользователя"
// @Success 200 {object} userservice.User "Получение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /user [get]
// @Security LogisticAuth
func GetUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := client.GetUser(c.Context(), c.Query("id"))
		if err != nil {
			return err
		}

		return c.JSON(user)
	}
}
