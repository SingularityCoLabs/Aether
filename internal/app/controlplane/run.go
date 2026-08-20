// Package controlplane composes aetherd's infrastructure and transport.
package controlplane

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/SingularityCoLabs/aether/internal/buildinfo"
	"github.com/SingularityCoLabs/aether/internal/config"
	"github.com/SingularityCoLabs/aether/internal/database"
	"github.com/SingularityCoLabs/aether/internal/server"
)

// Run starts aetherd and blocks until it fails or the context is canceled.
func Run(ctx context.Context, cfg config.Config, logger *slog.Logger) error {
	startedAt := time.Now().UTC()

	var connection *database.Database
	if cfg.DatabaseURL != "" {
		migrationCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		err := database.Migrate(migrationCtx, cfg.DatabaseURL)
		cancel()
		if err != nil {
			return err
		}

		openCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		connection, err = database.Open(openCtx, cfg.DatabaseURL)
		cancel()
		if err != nil {
			return err
		}
		defer connection.Close()
	} else {
		logger.Warn("database is disabled; aetherd is suitable only for Phase 0 development")
	}

	httpServer := &http.Server{
		Addr: cfg.HTTPAddress,
		Handler: server.NewHandler(server.Options{
			Logger:         logger,
			Build:          buildinfo.Current(),
			Environment:    cfg.Environment,
			StartedAt:      startedAt,
			Database:       connection,
			AllowedOrigins: cfg.AllowedOrigins,
		}),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       2 * time.Minute,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("aetherd listening",
			"address", cfg.HTTPAddress,
			"environment", cfg.Environment,
			"version", buildinfo.Version,
		)
		errCh <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return nil
	}
}
