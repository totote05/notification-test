package main

import (
	"notificator/builder"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	httpServerBuilder := builder.FileRegistryServerBuilder{}
	httpServer := httpServerBuilder.
		SetHost().
		SetUserRepository().
		SetNotificationBuilder().
		SetRegistryRepository().
		GetHttServer()

	go httpServer.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	httpServer.Stop()
}
