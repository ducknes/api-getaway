package authservice

type LoginUser struct {
	Username string `json:"username"` // Логин
	Password string `json:"password"` // Пароль
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`  // Токен пользователя
	RefreshToken string `json:"refresh_token"` // Токен для обновления
}
