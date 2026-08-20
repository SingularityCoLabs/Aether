package config

import (
	"log/slog"
	"testing"
	"time"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("AETHER_ENVIRONMENT", "")
	t.Setenv("AETHER_HTTP_ADDRESS", "")
	t.Setenv("AETHER_DATABASE_URL", "")
	t.Setenv("AETHER_DATABASE_REQUIRED", "")
	t.Setenv("AETHER_LOG_LEVEL", "")
	t.Setenv("AETHER_ALLOWED_ORIGINS", "")
	t.Setenv("AETHER_SHUTDOWN_TIMEOUT", "")

	got, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if got.Environment != "development" {
		t.Errorf("Environment = %q, want development", got.Environment)
	}
	if got.HTTPAddress != ":8080" {
		t.Errorf("HTTPAddress = %q, want :8080", got.HTTPAddress)
	}
	if got.LogLevel != slog.LevelInfo {
		t.Errorf("LogLevel = %v, want INFO", got.LogLevel)
	}
	if got.ShutdownTimeout != 10*time.Second {
		t.Errorf("ShutdownTimeout = %v, want 10s", got.ShutdownTimeout)
	}
}

func TestLoadRequiresDatabaseURL(t *testing.T) {
	t.Setenv("AETHER_DATABASE_REQUIRED", "true")
	t.Setenv("AETHER_DATABASE_URL", "")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want database validation error")
	}
}

func TestLoadRejectsInvalidOrigin(t *testing.T) {
	t.Setenv("AETHER_ALLOWED_ORIGINS", "localhost:3000/path")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want origin validation error")
	}
}
