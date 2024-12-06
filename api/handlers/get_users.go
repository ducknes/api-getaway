package handlers

import (
	"api-getaway/cluster/userservice"
	"github.com/gofiber/fiber/v2"
)

// GetUsersHandler получает пользователей.
//
// @Summary Получение пользователей
// @Description Получает определенное в limit количество пользователей
// @Tags Пользователи
// @Produce json
// @Param limit query int true "Количество пользователей"
// @Param cursor query string false "Ссылка на следующих пользователей"
// @Success 200 {object} []userservice.User "Получение прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизован"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /user/many [get]
// @Security LogisticAuth
func GetUsersHandler(client *userservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, cursor := parseQueryParams(c)

		users, err := client.GetUsers(c.Context(), limit, cursor)
		if err != nil {
			return err
		}

		return c.JSON(users)
	}
}
