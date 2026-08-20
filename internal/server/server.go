// Package server exposes Aether's HTTP and ConnectRPC transport.
package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"slices"
	"time"

	aetherv1connect "github.com/SingularityCoLabs/aether/gen/go/aether/v1/aetherv1connect"
	"github.com/SingularityCoLabs/aether/internal/buildinfo"
)

const readinessTimeout = 2 * time.Second

// HealthChecker is implemented by required runtime dependencies.
type HealthChecker interface {
	Ping(context.Context) error
}

// Options configures the public Phase 0 transport.
type Options struct {
	Logger         *slog.Logger
	Build          buildinfo.Info
	Environment    string
	StartedAt      time.Time
	Database       HealthChecker
	AllowedOrigins []string
}

// NewHandler builds a transport with health and the typed SystemService only.
func NewHandler(options Options) http.Handler {
	logger := options.Logger
	if logger == nil {
		logger = slog.Default()
	}
	if options.StartedAt.IsZero() {
		options.StartedAt = time.Now().UTC()
	}

	mux := http.NewServeMux()
	systemPath, systemHandler := aetherv1connect.NewSystemServiceHandler(&systemService{
		build:       options.Build,
		environment: options.Environment,
		startedAt:   options.StartedAt,
	})
	mux.Handle(systemPath, systemHandler)

	mux.HandleFunc("GET /livez", func(writer http.ResponseWriter, _ *http.Request) {
		writeJSON(writer, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /readyz", func(writer http.ResponseWriter, request *http.Request) {
		if options.Database == nil {
			writeJSON(writer, http.StatusOK, map[string]string{
				"status":   "ok",
				"database": "disabled",
			})
			return
		}
		ctx, cancel := context.WithTimeout(request.Context(), readinessTimeout)
		defer cancel()
		if err := options.Database.Ping(ctx); err != nil {
			logger.WarnContext(request.Context(), "readiness check failed",
				"component", "database",
				"error", err,
			)
			writeJSON(writer, http.StatusServiceUnavailable, map[string]string{
				"status":   "unavailable",
				"database": "unavailable",
			})
			return
		}
		writeJSON(writer, http.StatusOK, map[string]string{
			"status":   "ok",
			"database": "ok",
		})
	})
	mux.HandleFunc("GET /version", func(writer http.ResponseWriter, _ *http.Request) {
		writeJSON(writer, http.StatusOK, map[string]string{
			"name":        "aetherd",
			"version":     options.Build.Version,
			"commit":      options.Build.Commit,
			"buildDate":   options.Build.Date,
			"environment": options.Environment,
		})
	})

	handler := recoverPanics(logger, mux)
	return cors(options.AllowedOrigins, handler)
}

func writeJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Cache-Control", "no-store")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(value); err != nil {
		slog.Default().Error("write JSON response", "error", err)
	}
}

func cors(allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		origin := request.Header.Get("Origin")
		allowed := origin != "" && slices.Contains(allowedOrigins, origin)
		if allowed {
			writer.Header().Set("Access-Control-Allow-Origin", origin)
			writer.Header().Set("Access-Control-Allow-Credentials", "true")
			writer.Header().Add("Vary", "Origin")
		}
		if request.Method == http.MethodOptions {
			if !allowed {
				http.Error(writer, "origin not allowed", http.StatusForbidden)
				return
			}
			writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			writer.Header().Set(
				"Access-Control-Allow-Headers",
				"Authorization, Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms",
			)
			writer.Header().Set("Access-Control-Max-Age", "7200")
			writer.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func recoverPanics(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				logger.ErrorContext(request.Context(), "recovered HTTP panic", "panic", recovered)
				writeJSON(writer, http.StatusInternalServerError, map[string]string{
					"status": "internal_error",
				})
			}
		}()
		next.ServeHTTP(writer, request)
	})
}
