package mock

import (
	"notificator/core"
	"notificator/notification"
	"notificator/registry"
	"time"
)

func getChannelResponse(channel core.NotificationType) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_OK,
		Channel: channel,
	}
}

func NewRegistryResult(category core.Category, message string, registeredAt time.Time) []registry.Record {
	records := []registry.Record{}

	users := NewUserList()
	for _, user := range users {
		if core.Contains(user.Subscribed, category) {
			for _, channel := range user.Channels {
				records = append(records, registry.Record{
					User:         user,
					Category:     category,
					Response:     getChannelResponse(channel),
					Message:      message,
					RegisteredAt: registeredAt,
				})
			}
		}
	}

	return records
}
