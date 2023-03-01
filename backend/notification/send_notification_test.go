package notification_test

import (
	"notificator/builder"
	"notificator/core"
	"notificator/mock"
	"notificator/notification"
	"notificator/user"
	"testing"
)

func TestSuccessSendNotification(t *testing.T) {
	message := "Test message"

	repo := mock.NewUserRepository()
	model := user.NewUserModel(repo)

	notifierBuilder := builder.SuccessNotifierBuilder{}
	notifier := notifierBuilder.SetServices().GetNotifier()

	categories := []core.Category{core.MOVIES, core.FINANCE, core.SPORT}
	for _, category := range categories {
		for _, user := range model.GetUsersByCategory(category) {
			responses := notifier.SendNotification(user, category, message)

			if len(responses) != len(user.Channels) {
				t.Errorf("failed identificating services %v", user.Channels)
			}

			for _, response := range responses {
				if response.Status != notification.RESPONSE_OK || !core.Contains(user.Channels, response.Channel) {
					t.Errorf("failed on success sending service %v", response.Channel)
				}
			}
		}
	}
}

func TestFailedSendNotification(t *testing.T) {
	message := "Test message"

	repo := mock.NewUserRepository()
	model := user.NewUserModel(repo)

	notifierBuilder := builder.FailNotifierBuilder{}
	notifier := notifierBuilder.SetServices().GetNotifier()

	categories := []core.Category{core.MOVIES, core.FINANCE, core.SPORT}
	for _, category := range categories {
		for _, user := range model.GetUsersByCategory(category) {
			responses := notifier.SendNotification(user, category, message)

			if len(responses) != len(user.Channels) {
				t.Errorf("failed identificating services %v", user.Channels)
			}

			for _, response := range responses {
				if response.Status != notification.RESPONSE_FAIL || !core.Contains(user.Channels, response.Channel) {
					t.Errorf("failed on failure sending service: %v", response.Channel)
				}
			}
		}
	}
}
