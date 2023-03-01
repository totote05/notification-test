package registry_test

import (
	"notificator/core"
	"notificator/mock"
	"notificator/notification"
	"notificator/registry"
	"testing"
	"time"
)

func TestSaveLog(t *testing.T) {
	message := "Test message"
	user := mock.NewUserRepository().GetUsers()[0]
	category := core.MOVIES
	response := notification.Response{
		Status:  notification.RESPONSE_OK,
		Channel: core.EMAIL,
	}

	registryRepo := mock.NewRegistryRepository()
	registryModel := registry.NewModel(registryRepo)

	if repo, ok := registryRepo.(*mock.RegistryMockRepository); ok {
		repo.SetFailureMode(false)
	}

	log := registry.Record{
		User:         user,
		Category:     category,
		Response:     response,
		Message:      message,
		RegisteredAt: time.Now(),
	}

	if err := registryModel.SaveRecord(log); err != nil {
		t.Errorf("failed to save log: %s", err)
	}

	if repo, ok := registryRepo.(*mock.RegistryMockRepository); ok {
		repo.SetFailureMode(true)
	}

	if err := registryModel.SaveRecord(log); err == nil {
		t.Errorf("failed to save log: %s", err)
	}
}
