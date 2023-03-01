package mock

import "notificator/notification"

type SuccessServiceFactory struct{}

func (*SuccessServiceFactory) NewEmailService() notification.Service {
	return &EmailService{}
}

func (*SuccessServiceFactory) NewPushService() notification.Service {
	return &PushService{}
}

func (*SuccessServiceFactory) NewSmsService() notification.Service {
	return &SmsService{}
}

type FailedServiceFactory struct{}

func (*FailedServiceFactory) NewEmailService() notification.Service {
	return &FailEmailService{}
}

func (*FailedServiceFactory) NewPushService() notification.Service {
	return &FailPushService{}
}

func (*FailedServiceFactory) NewSmsService() notification.Service {
	return &FailSmsService{}
}

var _ notification.ServiceFactory = (*SuccessServiceFactory)(nil)
var _ notification.ServiceFactory = (*FailedServiceFactory)(nil)
