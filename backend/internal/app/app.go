package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kldd0/travel-hack-2024/internal/config"
	v1 "github.com/kldd0/travel-hack-2024/internal/controller/http/v1"
	httpserver "github.com/kldd0/travel-hack-2024/pkg/http-server"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Run() {
	config := config.MustLoad()

	SetLogger(config.Env)

	// Echo handler
	log.Info().Msg("Initializing handlers and routes...")
	handler := echo.New()

	// Router
	v1.NewRouter(handler)

	// HTTP server
	log.Info().Msg("Starting http server")
	log.Debug().Msgf("Server addr: %s", config.HTTPServer.Address)
	httpServer := httpserver.New(
		handler,
		httpserver.Addr(config.HTTPServer.Address),
		httpserver.ReadTimeout(config.HTTPServer.Timeout),
		httpserver.WriteTimeout(config.HTTPServer.Timeout),
	)

	// Waiting signal
	log.Info().Msg("Configuring graceful shutdown")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msgf("app.Run.signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Err(fmt.Errorf("app.Run.httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info().Msg("Shutting down")
	err := httpServer.Shutdown()
	if err != nil {
		log.Err(fmt.Errorf("app.Run.httpServer.Shutdown: %w", err))
	}
}
