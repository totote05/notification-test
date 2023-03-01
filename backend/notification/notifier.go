package notification

import (
	"fmt"
	"notificator/core"
	"notificator/user"
)

type Notifier struct {
	services []Service
}

func NewNotifier(services []Service) *Notifier {
	return &Notifier{
		services: services,
	}
}

func (n *Notifier) SendNotification(
	user user.User,
	category core.Category,
	message string,
) []Response {
	responses := []Response{}

	for _, service := range n.services {
		if core.Contains(user.Channels, service.GetType()) {
			contact := getContactByChannel(user, service.GetType())
			response := service.Send(user.Name, contact, message)
			responses = append(responses, response)
		}
	}

	return responses
}

func getContactByChannel(user user.User, channel core.NotificationType) core.Contact {
	switch channel {
	case core.EMAIL:
		return user.Email
	case core.PUSH, core.SMS:
		return user.PhoneNumber
	}
	panic(fmt.Errorf("unknown channel %s", channel))
}
