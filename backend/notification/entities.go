package notification

import "notificator/core"

type ResponseStatus string

const (
	RESPONSE_OK   ResponseStatus = "response_ok"
	RESPONSE_FAIL ResponseStatus = "response_fail"
)

type Response struct {
	Status  ResponseStatus        `json:"status"`
	Channel core.NotificationType `json:"channel"`
}
