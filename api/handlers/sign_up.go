package handlers

import (
	"api-getaway/cluster/authservice"
	"github.com/gofiber/fiber/v2"
)

// SignUpHandler Регистрация.
//
// @Summary Регистрация
// @Description регистрирует пользователя в сервисе
// @Tags Авторизация
// @Produce json
// @Param registration body authservice.LoginUser true "Данные пользователя"
// @Success 200 {object} authservice.LoginResponse "Регистрация прошла успешно"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /sign-up [post]
func SignUpHandler(authService *authservice.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registration authservice.LoginUser
		if err := c.BodyParser(&registration); err != nil {
			return err
		}

		token, err := authService.SignUp(c.UserContext(), registration)
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
