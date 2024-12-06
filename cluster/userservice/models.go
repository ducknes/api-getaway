package userservice

type Role uint

const (
	Undefined Role = iota
	UserRole
	AdminRole
)

type User struct {
	Id                string `json:"id,omitempty"`      // Id пользователя
	Username          string `json:"username"`          // Логин
	Surname           string `json:"surname"`           // Фамилия
	Name              string `json:"name"`              // Имя
	Lastname          string `json:"lastname"`          // Отчетство
	RegisteredObjects int    `json:"registeredObjects"` // Количестов подтвержленных продуктов
	Role              Role   `json:"role"`              // Роль
}
