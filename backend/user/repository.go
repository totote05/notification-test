package user

type Repository interface {
	GetUsers() []User
}
