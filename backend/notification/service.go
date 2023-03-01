package notification

import "notificator/core"

type Service interface {
	GetType() core.NotificationType
	Send(name string, contact core.Contact, message string) Response
}
