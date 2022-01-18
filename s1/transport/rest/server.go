package rest

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	Server *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	err := s.Server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server running: %w", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
