package user

import "notificator/core"

type User struct {
	Id          string                  `json:"id"`
	Name        string                  `json:"name"`
	Email       core.Contact            `json:"email"`
	PhoneNumber core.Contact            `json:"phoneNumber"`
	Subscribed  []core.Category         `json:"subscribed"`
	Channels    []core.NotificationType `json:"channels"`
}
