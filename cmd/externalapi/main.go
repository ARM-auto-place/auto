package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	accessClient "github.com/ARM-auto-place/access/pkg/client/transport"
	"github.com/ARM-auto-place/auto/internal/config"
	autoApiService "github.com/ARM-auto-place/auto/internal/service"
	"github.com/ARM-auto-place/auto/internal/storage/postgres"
	"github.com/ARM-auto-place/auto/internal/transport/jsonRPC/externalapi"
	"github.com/ARM-auto-place/auto/internal/transport/jsonRPC/middlewares"
	publicErrors "github.com/ARM-auto-place/auto/pkg/models/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

const serviceName = "auto-api"

func main() {
	cfg := config.Values()
	logger := cfg.Logger().With().Str("serviceName", serviceName).Logger()
	log.Logger = logger

	cm, err := postgres.New(logger, cfg.Postgres)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to postgres")
	}
	defer cm.GetMasterConn().Close()
	defer cm.GetReplicaConn().Close()

	storage := postgres.NewAutoStorage(logger, cm)
	access := accessClient.New(
		cfg.AccessAddr,
		accessClient.DecodeError(publicErrors.ErrorDecoder),
	).InternalAPI()

	svc := autoApiService.NewPayAPIService(logger, storage)
	svc = middlewares.NewInternalMiddleware(svc, access)

	app := externalapi.New(
		logger,
		externalapi.AutoApi(externalapi.NewAutoApi(svc)),
		externalapi.MaxBodySize(cfg.MaxRequestBodySize),
		externalapi.ReadTimeout(cfg.ReadTimeout),
		externalapi.WriteTimeout(cfg.WriteTimeout),
	).WithLog().WithMetrics()

	server := &fasthttp.Server{
		Handler:            app.Fiber().Handler(),
		MaxRequestBodySize: cfg.MaxRequestBodySize,
		ReadBufferSize:     cfg.MaxRequestHeaderSize,
		ReadTimeout:        cfg.ReadTimeout,
		WriteTimeout:       cfg.WriteTimeout,
	}

	app.ServeHealth(cfg.HealthBind, fiber.Map{"status": "ok"})
	app.ServeMetrics(logger, cfg.MetricsPath, cfg.MetricsBind)

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- server.ListenAndServe(cfg.ServiceBind)
	}()

	logger.Info().Str("address", cfg.ServiceBind).Msg("external API server started")
	select {
	case <-shutdownCtx.Done():
		logger.Info().Msg("shutdown signal received")
	case err := <-serveErr:
		if err != nil {
			logger.Error().Err(err).Msg("external API server stopped")
		} else {
			logger.Info().Msg("external API server stopped")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := server.ShutdownWithContext(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to shutdown external API server")
	}
	app.Shutdown()
	logger.Info().Msg("external API server shutdown complete")
}
