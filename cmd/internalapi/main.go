package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ARM-auto-place/auto/internal/config"
	autoApiService "github.com/ARM-auto-place/auto/internal/service"
	"github.com/ARM-auto-place/auto/internal/storage/postgres"
	"github.com/ARM-auto-place/auto/internal/transport/jsonRPC/internalapi"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

const serviceName = "auto-api-internal"

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
	svc := autoApiService.NewInternalPayAPIService(logger, storage)

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := internalapi.New(
		logger,
		internalapi.SetReadBufferSize(cfg.MaxRequestHeaderSize),
		internalapi.MaxBodySize(cfg.MaxRequestBodySize),
		internalapi.ReadTimeout(cfg.ReadTimeout),
		internalapi.WriteTimeout(cfg.WriteTimeout),
		internalapi.MaxBatchSize(100),
		internalapi.MaxBatchWorkers(1),
		internalapi.WithRequestID("X-Request-Id"),
		internalapi.Use(func(ctx *fiber.Ctx) error {
			requestCtx, cancel := context.WithTimeout(shutdownCtx, cfg.WriteTimeout)
			defer cancel()
			ctx.SetUserContext(log.Ctx(ctx.UserContext()).WithContext(requestCtx))
			return ctx.Next()
		}),
		internalapi.InternalAutoAPI(internalapi.NewInternalAutoAPI(svc)),
	).WithLog().WithMetrics()

	app.ServeHealth(cfg.HealthBind, fiber.Map{"status": "ok"})
	app.ServeMetrics(logger, cfg.MetricsPath, cfg.MetricsBind)

	serveErr := make(chan error, 1)
	go func() {
		logger.Info().Str("address", cfg.ServiceBind).Msg("internal API server started")
		serveErr <- app.Fiber().Listen(cfg.ServiceBind)
	}()

	select {
	case <-shutdownCtx.Done():
		logger.Info().Msg("shutdown signal received")
	case err := <-serveErr:
		if err != nil {
			logger.Error().Err(err).Msg("internal API server stopped")
		} else {
			logger.Info().Msg("internal API server stopped")
		}
		stop()
	}

	done := make(chan struct{})
	go func() {
		app.Shutdown()
		close(done)
	}()

	select {
	case <-done:
		logger.Info().Msg("internal API server shutdown complete")
	case <-time.After(cfg.ShutdownTimeout):
		logger.Warn().Msg("internal API server shutdown timed out")
	}
}
