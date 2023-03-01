package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"notificator/core"
	"notificator/notification"
	"notificator/registry"
	"notificator/user"
	"time"
)

type MessagePayload struct {
	Category string `json:"category"`
	Message  string `json:"message"`
}

type MessageController struct {
	user     user.Model
	notifier *notification.Notifier
	registry *registry.Model
}

func NewMessageController(
	user user.Model,
	notifier *notification.Notifier,
	registry *registry.Model,
) *MessageController {
	return &MessageController{
		user:     user,
		notifier: notifier,
		registry: registry,
	}
}

func (c *MessageController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
	} else if r.Method == "GET" {
		if data, err := json.Marshal(c.GetAll()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	} else if r.Method == "POST" {
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("empty payload"))
			return
		}
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			return
		}

		payload := MessagePayload{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		}

		category, err := core.CategoryFromString(payload.Category)
		if payload.Message == "" || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid payload"))
		} else if err := c.Send(category, payload.Message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *MessageController) Send(category core.Category, message string) error {
	for _, user := range c.user.GetUsersByCategory(category) {
		responses := c.notifier.SendNotification(
			user,
			category,
			message,
		)

		for _, response := range responses {
			record := registry.Record{
				User:         user,
				Category:     category,
				Response:     response,
				Message:      message,
				RegisteredAt: time.Now().Truncate(time.Microsecond),
			}

			if err := c.registry.SaveRecord(record); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *MessageController) GetAll() []registry.Record {
	return c.registry.GetRecords()
}
