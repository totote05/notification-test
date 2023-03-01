package mock

import (
	"notificator/user"
)

type UserMockRepository struct {
	users []user.User
}

func NewUserRepository() user.Repository {
	return &UserMockRepository{
		users: NewUserList(),
	}
}

func (r *UserMockRepository) GetUsers() []user.User {
	return r.users
}

var _ user.Repository = (*UserMockRepository)(nil)
