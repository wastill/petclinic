package biz

type User struct {
	Username string
	Password string
	Salt     string
	Enable   bool
	Roles    []string
}
type CreateUserParam struct {
	Username string
	Password string
	Enable   bool
}

func NewCreateUserRequest() {

}
