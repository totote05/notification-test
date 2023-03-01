package core

type Contact string

type Category string

const (
	SPORT   Category = "sport"
	FINANCE Category = "finance"
	MOVIES  Category = "movies"
)

type NotificationType string

const (
	SMS   NotificationType = "sms"
	EMAIL NotificationType = "email"
	PUSH  NotificationType = "push"
)
