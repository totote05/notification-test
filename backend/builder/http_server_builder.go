package builder

import (
	"notificator/mock"
	"notificator/registry"
	"notificator/server"
	"notificator/user"
)

type HttpServerBuilder interface {
	SetHost() HttpServerBuilder
	SetUserRepository() HttpServerBuilder
	SetRegistryRepository() HttpServerBuilder
	SetNotificationBuilder() HttpServerBuilder
	GetHttServer() *server.HttpServer
}

type FileRegistryServerBuilder struct {
	host     string
	user     user.Repository
	registry registry.Repository
	notifier NotifierBuilder
}

func (b *FileRegistryServerBuilder) GetHttServer() *server.HttpServer {

	controller := server.NewMessageController(
		user.NewUserModel(b.user),
		b.notifier.SetServices().GetNotifier(),
		registry.NewModel(b.registry),
	)

	httpServer := server.NewHttpServer(b.host)
	httpServer.RegisterController("/message", controller)

	return httpServer
}

func (b *FileRegistryServerBuilder) SetHost() HttpServerBuilder {
	b.host = ":5000"
	return b
}

func (b *FileRegistryServerBuilder) SetNotificationBuilder() HttpServerBuilder {
	b.notifier = &SuccessNotifierBuilder{}
	return b
}

func (b *FileRegistryServerBuilder) SetRegistryRepository() HttpServerBuilder {
	b.registry = registry.NewFileRepository("data.json")
	return b
}

func (b *FileRegistryServerBuilder) SetUserRepository() HttpServerBuilder {
	b.user = mock.NewUserRepository()
	return b
}

var _ HttpServerBuilder = (*FileRegistryServerBuilder)(nil)
