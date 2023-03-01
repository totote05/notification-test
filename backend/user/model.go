package user

import "notificator/core"

type Model struct {
	repo Repository
}

func NewUserModel(repo Repository) Model {
	return Model{repo}
}

func (m *Model) GetUsersByCategory(category core.Category) []User {
	users := []User{}

	for _, user := range m.repo.GetUsers() {
		if core.Contains(user.Subscribed, category) {
			users = append(users, user)
		}
	}

	return users
}
