package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

// UpdateUserHandler обновляет пользователя.
//
// @Summary Обновление пользователя
// @Description Обновляет пользователя
// @Tags Пользователи
// @Param user body userservice.User true "Обновляемый пользователь"
// @Success 200 "Обновление прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /user [put]
// @Security LogisticAuth
func UpdateUserHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userservice.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		return client.UpdateUser(c.Context(), user)
	}
}
