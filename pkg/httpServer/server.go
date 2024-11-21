package httpServer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"syscall"
	"time"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(config *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    config.HTTPServerAddress,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		fmt.Println("caught signal", map[string]string{
			"signal": sig.String()})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- s.httpServer.Shutdown(ctx)

	}()

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdownError; err != nil {
		return err
	}

	slog.Info("stopped server", slog.String("addr", s.httpServer.Addr))

	return nil
}
