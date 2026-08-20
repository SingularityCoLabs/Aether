// Package observability configures telemetry shared by Aether processes.
package observability

import (
	"io"
	"log/slog"
)

// NewLogger returns a structured JSON logger suitable for local and container
// environments. Sensitive values must be redacted before they reach it.
func NewLogger(output io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level: level,
	}))
}
