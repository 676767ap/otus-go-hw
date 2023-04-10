package internalhttp

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	logger Logger
	app    Application
	server *http.Server
}

type Logger interface { // TODO
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application) *Server {
	s := &http.Server{
		Addr:              ":8087",
		Handler:           &handler{},
		ReadHeaderTimeout: 10 * time.Second,
	}

	return &Server{
		app:    app,
		logger: logger,
		server: s,
	}
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
