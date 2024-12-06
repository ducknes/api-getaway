package handlers

import (
	"api-getaway/cluster/authservice"
	"github.com/gofiber/fiber/v2"
)

// RefreshHandler Обновление токена.
//
// @Summary Обновление токена
// @Description обновляет токен авторизации по токену обновления
// @Tags Авторизация
// @Success 200 {object} authservice.LoginResponse "Обновление токена прошло успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /refresh [post]
func RefreshHandler(authService *authservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := authService.RefreshToken(c.UserContext(), c.Cookies("refresh_token"))
		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:  "refresh_token",
			Value: token.RefreshToken,
		})

		return c.JSON(fiber.Map{
			"access_token": token.AccessToken,
		})
	}
}
