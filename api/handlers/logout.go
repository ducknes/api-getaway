package handlers

import (
	"api-getaway/cluster/authservice"
	"github.com/gofiber/fiber/v2"
)

// LogoutHandler выход.
//
// @Summary Выход
// @Description разлогинивает пользователя с сайта
// @Tags Авторизация
// @Success 200 "Выход прошел успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /logout [post]
func LogoutHandler(authService *authservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "refresh_token",
			Value: "dead",
		})

		return authService.Logout(c.UserContext(), c.Cookies("refresh_token"))
	}
}
