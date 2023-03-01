package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"notificator/builder"
	"notificator/mock"
	"notificator/registry"
	"notificator/server"
	"notificator/user"
	"os"
	"testing"
)

const test_file = "data-test.json"
const test_host = ":3000"
const test_endpoint = "http://localhost:3000/message"

func beforeEach() (FileRegistryServerBuilder, *server.HttpServer) {
	httpServerBuilder := FileRegistryServerBuilder{}
	httpServer := httpServerBuilder.
		SetHost().
		SetUserRepository().
		SetNotificationBuilder().
		SetRegistryRepository().
		GetHttServer()

	return httpServerBuilder, httpServer
}

func TestGetMessage(t *testing.T) {
	os.Remove(test_file)

	httpServerBuilder, httpServer := beforeEach()
	defer httpServer.Stop()

	go httpServer.Start()

	request, err := http.NewRequest("GET", test_endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("fail get request, wrong status code: got %v want %v", status, http.StatusOK)
	}

	records := []registry.Record{}
	if err := json.Unmarshal(responseRecorder.Body.Bytes(), &records); err != nil {
		t.Fatal(err)
	}

	expected := 0
	result := len(records)
	if result != expected {
		t.Errorf("fail get request, unexpected body: got %d want %d", result, expected)
	}
}

func TestPostInvalidPayloadMessage(t *testing.T) {
	os.Remove(test_file)

	expected := "invalid payload"
	testData := []struct {
		payload  string
		expected string
		message  string
	}{
		{"", "empty payload", "empty payload"},
		{"{}", "invalid payload", "empty payload"},
		{`{"category": ""}`, "invalid payload", "empty category"},
		{`{"message": ""}`, "invalid payload", "empty message"},
		{`{"category": "", "message": ""}`, "invalid payload", "empty category and message"},
		{`{"category": "", "message": "test"}`, "invalid payload", "empty category and not empty message"},
		{`{"category": "movies", "message": ""}`, "invalid payload", "not empty category and empty message"},
		{`{"category": "movie", "message": ""}`, "invalid payload", "wrong category and empty message"},
		{`{"category": "movie", "message": "test"}`, "invalid payload", "wrong category and not empty message"},
	}

	httpServerBuilder, httpServer := beforeEach()
	defer httpServer.Stop()

	go httpServer.Start()

	for _, data := range testData {
		var payload io.Reader

		if data.payload != "" {
			payload = bytes.NewBuffer([]byte(data.payload))
		}

		request, err := http.NewRequest("POST", test_endpoint, payload)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != http.StatusBadRequest {
			t.Errorf("fail post request with %s, wrong status code: got %v want %v", data.message, status, http.StatusOK)
		}

		if body := responseRecorder.Body.String(); body != data.expected {
			t.Errorf("fail post request with %s, wrong message: got '%v' want '%v'", data.message, body, expected)
		}
	}
}

func TestPostValidPayloadMessage(t *testing.T) {
	os.Remove(test_file)

	expected := "invalid payload"
	testData := []struct {
		payload  string
		expected string
		message  string
	}{
		{`{"category": "sport", "message": "test sport"}`, "", "sport"},
		{`{"category": "movies", "message": "test movies"}`, "", "movies"},
		{`{"category": "finance", "message": "test finance"}`, "", "finance"},
	}

	httpServerBuilder, httpServer := beforeEach()
	defer httpServer.Stop()

	go httpServer.Start()

	for _, data := range testData {
		request, err := http.NewRequest("POST", test_endpoint, bytes.NewBuffer([]byte(data.payload)))
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != http.StatusCreated {
			t.Errorf("fail post request with %s, wrong status code: got %v want %v", data.message, status, http.StatusOK)
		}

		if body := responseRecorder.Body.String(); body != data.expected {
			t.Errorf("fail post request with %s, wrong message: got '%v' want '%v'", data.message, body, expected)
		}
	}
}

func TestIntegrationMessage(t *testing.T) {
	os.Remove(test_file)

	// empty registry
	httpServerBuilder, httpServer := beforeEach()
	defer httpServer.Stop()

	go httpServer.Start()

	request, err := http.NewRequest("GET", test_endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("fail get request, wrong status code: got %v want %v", status, http.StatusOK)
	}

	records := []registry.Record{}
	if err := json.Unmarshal(responseRecorder.Body.Bytes(), &records); err != nil {
		t.Fatal(err)
	}

	expected := 0
	result := len(records)
	if result != expected {
		t.Errorf("fail get request, unexpected body: got %d want %d", result, expected)
	}

	// post a message
	payload := bytes.NewBuffer([]byte(`{"category": "movies", "message": "test movies"}`))
	request, err = http.NewRequest("POST", test_endpoint, payload)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder = httptest.NewRecorder()
	httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("fail post request, wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := responseRecorder.Body.String(); body != "" {
		t.Errorf("fail post request, wrong message: got '%v' want '%v'", body, expected)
	}

	// filled registry
	request, err = http.NewRequest("GET", test_endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder = httptest.NewRecorder()
	httpServerBuilder.controller.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("fail get request, wrong status code: got %v want %v", status, http.StatusOK)
	}

	records = []registry.Record{}
	if err := json.Unmarshal(responseRecorder.Body.Bytes(), &records); err != nil {
		t.Fatal(err)
	}

	result = len(records)
	if result == 0 {
		t.Error("fail get request, registry is empty")
	}
}

type FileRegistryServerBuilder struct {
	host       string
	user       user.Repository
	registry   registry.Repository
	notifier   builder.NotifierBuilder
	controller *server.MessageController
}

func (b *FileRegistryServerBuilder) GetHttServer() *server.HttpServer {
	b.controller = server.NewMessageController(
		user.NewUserModel(b.user),
		b.notifier.SetServices().GetNotifier(),
		registry.NewModel(b.registry),
	)

	httpServer := server.NewHttpServer(b.host)
	httpServer.RegisterController("/message", b.controller)

	return httpServer
}

func (b *FileRegistryServerBuilder) SetHost() builder.HttpServerBuilder {
	b.host = test_host
	return b
}

func (b *FileRegistryServerBuilder) SetNotificationBuilder() builder.HttpServerBuilder {
	b.notifier = &builder.SuccessNotifierBuilder{}
	return b
}

func (b *FileRegistryServerBuilder) SetRegistryRepository() builder.HttpServerBuilder {
	b.registry = registry.NewFileRepository(test_file)
	return b
}

func (b *FileRegistryServerBuilder) SetUserRepository() builder.HttpServerBuilder {
	b.user = mock.NewUserRepository()
	return b
}

var _ builder.HttpServerBuilder = (*FileRegistryServerBuilder)(nil)
