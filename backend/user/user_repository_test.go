package user_test

import (
	"notificator/mock"
	"notificator/user"
	"reflect"
	"testing"
)

func compareUsers(from, to []user.User) bool {
	return reflect.DeepEqual(from, to)
}

func TestUserRepository(t *testing.T) {
	repository := mock.NewUserRepository()

	if !compareUsers(mock.NewUserList(), repository.GetUsers()) {
		t.Error("user repository failed: the mocked users are not the same on the repository")
	}
}
