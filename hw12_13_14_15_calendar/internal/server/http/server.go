package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/676767ap/otus-go-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/676767ap/otus-go-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/pkg/errors"
)

type Server struct {
	app    *app.App
	server *http.Server
}

func NewServer(app *app.App, config *config.Config) *Server {
	handler := decorateHandler(app.GetLogger())

	return &Server{
		app: app,
		server: &http.Server{
			Addr:              net.JoinHostPort(config.HTTP.Host, config.HTTP.Port),
			Handler:           handler,
			ReadTimeout:       config.HTTP.ReadTimeout,
			WriteTimeout:      config.HTTP.WriteTimeout,
			ReadHeaderTimeout: config.HTTP.ReadHeaderTimeout,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.app.GetLogger().Info("Start server...")

	err := s.server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("can't start server, %w", err)
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.app.GetLogger().Info("Stop server...")

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("can't shutdown server, %w", err)
	}
	return nil
}

func decorateHandler(logger app.Logger) http.HandlerFunc {
	handler := http.HandlerFunc(testAction)
	handler = headersMiddleware(handler)
	handler = loggingMiddleware(handler, logger)

	return handler
}

func testAction(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("hello-world")
	if err != nil {
		return
	}
}
