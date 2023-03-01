package mock

import (
	"notificator/core"
	"notificator/user"
)

var user1 = user.User{
	Id:          "4dJpnhNhme63ziAHUCXtq",
	Name:        "Emily Johnson",
	Email:       "emily.johnson@fakemail.com",
	PhoneNumber: "+1 415 4459321",
	Subscribed:  []core.Category{core.SPORT, core.MOVIES},
	Channels:    []core.NotificationType{core.PUSH},
}
var user2 = user.User{
	Id:          "a4fw2H7BKTTYDfr3CHkGj",
	Name:        "Mateo Gonzalez",
	Email:       "mateo.gonzalez@fakemail.com",
	PhoneNumber: "+54 343 4158693",
	Subscribed:  []core.Category{core.SPORT, core.FINANCE},
	Channels:    []core.NotificationType{core.EMAIL, core.SMS, core.PUSH},
}
var user3 = user.User{
	Id:          "HR6jRkpcLw7eEm77g4wE2",
	Name:        "Michael Davis",
	Email:       "michael.davis@fakemail.com",
	PhoneNumber: "+1 212 4547962",
	Subscribed:  []core.Category{core.MOVIES},
	Channels:    []core.NotificationType{core.EMAIL, core.SMS},
}
var user4 = user.User{
	Id:          "dFW6gEHmmm92gCALcLNVD",
	Name:        "Valentina Gomez",
	Email:       "valentina.gomez@fakemail.com",
	PhoneNumber: "+54 343 4693214",
	Subscribed:  []core.Category{core.SPORT, core.FINANCE, core.MOVIES},
	Channels:    []core.NotificationType{core.EMAIL},
}

func NewUserList() []user.User {
	return []user.User{user1, user2, user3, user4}
}
