package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kldd0/travel-hack-2024/internal/config"
	v1 "github.com/kldd0/travel-hack-2024/internal/controller/http/v1"
	httpserver "github.com/kldd0/travel-hack-2024/internal/pkg/http-server"
	"github.com/kldd0/travel-hack-2024/internal/pkg/validator"
	"github.com/kldd0/travel-hack-2024/internal/repository"
	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@title			Tour Management Service
//	@version		1.0
//	@description	Additional service for RUSSPASS

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization
//	@description				JWT token

func Run() {
	config := config.MustLoad()

	// Setup logger
	SetLogger(config.Env)

	/*
		// Repositories
		log.Info().Msg("Initializing postgres...")
		pg, err := postgres.New(config.Postgres.DSN, postgres.MaxPoolSize(config.Postgres.MaxPoolSize))
		if err != nil {
			log.Err(fmt.Errorf("app.Run.postgres.NewServices: %w", err))
			os.Exit(-1)
		}
		defer pg.Close()
	*/

	// Repositories
	log.Info().Msg("Initializing repositories")
	repositories := repository.NewRepositories()

	// Services dependencies
	log.Info().Msg("Initializing services")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	// Echo handler
	log.Info().Msg("Initializing handlers and routes")
	handler := echo.New()
	// setup handler validator as lib validator
	handler.Validator = validator.NewCustomValidator()

	// Router
	v1.NewRouter(handler, services)

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
	if err := httpServer.Shutdown(); err != nil {
		log.Err(fmt.Errorf("app.Run.httpServer.Shutdown: %w", err))
	}
}
