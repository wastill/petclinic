package biz

type User struct {
	Username string
	Password string
	Enable   bool
}
type CreateUserParam struct {
	Username string
	Password string
	Enable   bool
}
