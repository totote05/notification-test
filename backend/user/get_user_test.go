package user_test

import (
	"notificator/core"
	"notificator/mock"
	"notificator/user"
	"testing"
)

func TestGetUser(t *testing.T) {
	repository := mock.NewUserRepository()
	model := user.NewUserModel(repository)

	users := model.GetUsersByCategory(core.SPORT)
	count := len(users)
	expected := 3
	if count != expected {
		t.Errorf("Fail on getting user by Sport category: expected %d, got %d", expected, count)
	}

	users = model.GetUsersByCategory(core.FINANCE)
	count = len(users)
	expected = 2
	if count != expected {
		t.Errorf("Fail on getting user by Finance category: expected %d, got %d", expected, count)
	}

	users = model.GetUsersByCategory(core.MOVIES)
	count = len(users)
	expected = 3
	if count != expected {
		t.Errorf("Fail on getting user by Movies category: expected %d, got %d", expected, count)
	}
}
