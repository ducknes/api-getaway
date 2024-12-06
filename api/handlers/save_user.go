package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

// SaveUserHandler сохраняет нового пользователя.
//
// @Summary Сохранение пользователя
// @Description Сохраняет нового пользователя
// @Tags Пользователи
// @Param user body userservice.User true "Сохраняемый пользователь"
// @Success 200 "Сохранение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /user [post]
// @Security LogisticAuth
func SaveUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userservice.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		return client.SaveUser(c.Context(), user)
	}
}
