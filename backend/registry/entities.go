package registry

import (
	"notificator/core"
	"notificator/notification"
	"notificator/user"
	"time"
)

type Record struct {
	User         user.User             `json:"user"`
	Category     core.Category         `json:"category"`
	Response     notification.Response `json:"response"`
	Message      string                `json:"message"`
	RegisteredAt time.Time             `json:"registeredAt"`
}

type SortByRegisteredAt []Record

func (a SortByRegisteredAt) Len() int {
	return len(a)
}

func (a SortByRegisteredAt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SortByRegisteredAt) Less(i, j int) bool {
	return a[i].RegisteredAt.After(a[j].RegisteredAt)
}
