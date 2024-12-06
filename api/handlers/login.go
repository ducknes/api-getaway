package handlers

import (
	"api-getaway/cluster/authservice"
	"github.com/gofiber/fiber/v2"
)

// LoginHandler Вход.
//
// @Summary Вход
// @Description авторизует пользователя на сайт и возвращает токен
// @Tags Авторизация
// @Produce json
// @Param login body authservice.LoginUser true "Данные пользователя"
// @Success 200 {object} authservice.LoginResponse "Авторизация прошла успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /login [post]
func LoginHandler(authService *authservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var login authservice.LoginUser
		if err := c.BodyParser(&login); err != nil {
			return err
		}

		token, err := authService.Login(c.UserContext(), login)
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
