// Package config loads and validates control-plane configuration.
package config

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultHTTPAddress    = ":8080"
	defaultEnvironment    = "development"
	defaultShutdown       = 10 * time.Second
	defaultAllowedOrigins = "http://localhost:3000"
)

// Config is the complete runtime configuration for aetherd.
type Config struct {
	Environment      string
	HTTPAddress      string
	DatabaseURL      string
	DatabaseRequired bool
	LogLevel         slog.Level
	AllowedOrigins   []string
	ShutdownTimeout  time.Duration
}

// Load reads configuration from environment variables and validates it.
func Load() (Config, error) {
	level, err := parseLogLevel(env("AETHER_LOG_LEVEL", "info"))
	if err != nil {
		return Config{}, err
	}

	databaseRequired, err := strconv.ParseBool(env("AETHER_DATABASE_REQUIRED", "false"))
	if err != nil {
		return Config{}, fmt.Errorf("parse AETHER_DATABASE_REQUIRED: %w", err)
	}

	shutdownTimeout, err := time.ParseDuration(env("AETHER_SHUTDOWN_TIMEOUT", defaultShutdown.String()))
	if err != nil {
		return Config{}, fmt.Errorf("parse AETHER_SHUTDOWN_TIMEOUT: %w", err)
	}

	cfg := Config{
		Environment:      env("AETHER_ENVIRONMENT", defaultEnvironment),
		HTTPAddress:      env("AETHER_HTTP_ADDRESS", defaultHTTPAddress),
		DatabaseURL:      strings.TrimSpace(os.Getenv("AETHER_DATABASE_URL")),
		DatabaseRequired: databaseRequired,
		LogLevel:         level,
		AllowedOrigins:   splitCSV(env("AETHER_ALLOWED_ORIGINS", defaultAllowedOrigins)),
		ShutdownTimeout:  shutdownTimeout,
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// Validate rejects configurations that could make health reporting misleading.
func (c Config) Validate() error {
	if strings.TrimSpace(c.Environment) == "" {
		return errors.New("AETHER_ENVIRONMENT must not be empty")
	}
	if _, err := net.ResolveTCPAddr("tcp", c.HTTPAddress); err != nil {
		return fmt.Errorf("invalid AETHER_HTTP_ADDRESS: %w", err)
	}
	if c.DatabaseRequired && c.DatabaseURL == "" {
		return errors.New("AETHER_DATABASE_URL is required when AETHER_DATABASE_REQUIRED=true")
	}
	if c.DatabaseURL != "" {
		parsed, err := url.Parse(c.DatabaseURL)
		if err != nil || (parsed.Scheme != "postgres" && parsed.Scheme != "postgresql") {
			return errors.New("AETHER_DATABASE_URL must be a postgres or postgresql URL")
		}
	}
	if c.ShutdownTimeout <= 0 {
		return errors.New("AETHER_SHUTDOWN_TIMEOUT must be greater than zero")
	}
	for _, origin := range c.AllowedOrigins {
		parsed, err := url.Parse(origin)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return fmt.Errorf("invalid allowed origin %q", origin)
		}
		if parsed.Path != "" {
			return fmt.Errorf("allowed origin %q must not contain a path", origin)
		}
	}
	return nil
}

func env(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func parseLogLevel(value string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(strings.ToUpper(strings.TrimSpace(value)))); err != nil {
		return 0, fmt.Errorf("parse AETHER_LOG_LEVEL: %w", err)
	}
	return level, nil
}

func splitCSV(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
