package mock

import (
	"notificator/core"
	"notificator/notification"
)

type SmsService struct{}

func (*SmsService) GetType() core.NotificationType {
	return core.SMS
}

func (s *SmsService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_OK,
		Channel: s.GetType(),
	}
}

type EmailService struct{}

func (*EmailService) GetType() core.NotificationType {
	return core.EMAIL
}

func (s *EmailService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_OK,
		Channel: s.GetType(),
	}
}

type PushService struct{}

func (*PushService) GetType() core.NotificationType {
	return core.PUSH
}

func (s *PushService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_OK,
		Channel: s.GetType(),
	}
}

type FailSmsService struct{}

func (*FailSmsService) GetType() core.NotificationType {
	return core.SMS
}

func (s *FailSmsService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_FAIL,
		Channel: s.GetType(),
	}
}

type FailEmailService struct{}

func (*FailEmailService) GetType() core.NotificationType {
	return core.EMAIL
}

func (s *FailEmailService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_FAIL,
		Channel: s.GetType(),
	}
}

type FailPushService struct{}

func (*FailPushService) GetType() core.NotificationType {
	return core.PUSH
}

func (s *FailPushService) Send(name string, contact core.Contact, message string) notification.Response {
	return notification.Response{
		Status:  notification.RESPONSE_FAIL,
		Channel: s.GetType(),
	}
}

var _ notification.Service = (*SmsService)(nil)
var _ notification.Service = (*EmailService)(nil)
var _ notification.Service = (*PushService)(nil)
var _ notification.Service = (*FailSmsService)(nil)
var _ notification.Service = (*FailEmailService)(nil)
var _ notification.Service = (*FailPushService)(nil)
