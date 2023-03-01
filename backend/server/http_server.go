package server

import (
	"context"
	"log"
	"net/http"
)

type HttpServer struct {
	server *http.Server
	mux    *http.ServeMux
}

func NewHttpServer(host string) *HttpServer {
	mux := http.NewServeMux()

	return &HttpServer{
		server: &http.Server{
			Addr:    host,
			Handler: mux,
		},
		mux: mux,
	}
}

func (s *HttpServer) RegisterController(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

func (s *HttpServer) Start() {
	log.Println("Server started on", s.server.Addr)
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server stopped incorrectly", err)
	} else {
		log.Println("Server stopped successfuly")
	}
}

func (s *HttpServer) Stop() {
	log.Print("Stopping server")
	s.server.Shutdown(context.Background())
}
