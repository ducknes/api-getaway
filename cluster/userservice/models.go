package userservice

type Role uint

const (
	Undefined Role = iota
	UserRole
	AdminRole
)

type User struct {
	Id                string `json:"id,omitempty"`
	Surname           string `json:"surname"`
	Name              string `json:"name"`
	Lastname          string `json:"lastname"`
	RegisteredObjects int    `json:"registeredObjects"`
	Role              Role   `json:"role"`
}
