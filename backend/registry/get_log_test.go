package registry_test

import (
	"notificator/builder"
	"notificator/core"
	"notificator/mock"
	"notificator/registry"
	"notificator/user"
	"testing"
	"time"
)

func TestGetLog(t *testing.T) {
	message := "Test message"
	registeredAt := time.Now()
	userModel := user.NewUserModel(mock.NewUserRepository())

	notifierBuilder := builder.SuccessNotifierBuilder{}
	notifier := notifierBuilder.SetServices().GetNotifier()

	categories := []core.Category{core.MOVIES, core.FINANCE, core.SPORT}
	for _, category := range categories {
		registryRepo := mock.NewRegistryRepository()
		registryModel := registry.NewModel(registryRepo)

		for _, user := range userModel.GetUsersByCategory(category) {
			responses := notifier.SendNotification(user, category, message)

			for _, response := range responses {
				log := registry.Record{
					User:         user,
					Category:     category,
					Response:     response,
					Message:      message,
					RegisteredAt: registeredAt,
				}

				if err := registryModel.SaveRecord(log); err != nil {
					t.Errorf("failed to save log: %s", err)
				}
			}
		}

		expected := len(mock.NewRegistryResult(category, message, registeredAt))
		results := len(registryModel.GetRecords())

		if expected != results {
			t.Errorf("failed getting logs, expected %d got %d", expected, results)
		}
	}
}
