package notification

type ServiceFactory interface {
	NewSmsService() Service
	NewEmailService() Service
	NewPushService() Service
}
