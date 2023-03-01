package builder

import (
	"notificator/mock"
	"notificator/notification"
)

type NotifierBuilder interface {
	SetServices() NotifierBuilder
	GetNotifier() *notification.Notifier
}

type SuccessNotifierBuilder struct {
	services []notification.Service
}

func (b *SuccessNotifierBuilder) SetServices() NotifierBuilder {
	serviceFactory := mock.SuccessServiceFactory{}
	b.services = []notification.Service{
		serviceFactory.NewSmsService(),
		serviceFactory.NewEmailService(),
		serviceFactory.NewPushService(),
	}

	return b
}

func (b *SuccessNotifierBuilder) GetNotifier() *notification.Notifier {
	return notification.NewNotifier(b.services)
}

type FailNotifierBuilder struct {
	services []notification.Service
}

func (b *FailNotifierBuilder) SetServices() NotifierBuilder {
	serviceFactory := mock.FailedServiceFactory{}
	b.services = []notification.Service{
		serviceFactory.NewSmsService(),
		serviceFactory.NewEmailService(),
		serviceFactory.NewPushService(),
	}

	return b
}

func (b *FailNotifierBuilder) GetNotifier() *notification.Notifier {
	return notification.NewNotifier(b.services)

}

var _ NotifierBuilder = (*SuccessNotifierBuilder)(nil)
var _ NotifierBuilder = (*FailNotifierBuilder)(nil)
