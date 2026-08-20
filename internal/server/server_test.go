package server

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	aetherv1 "github.com/SingularityCoLabs/aether/gen/go/aether/v1"
	aetherv1connect "github.com/SingularityCoLabs/aether/gen/go/aether/v1/aetherv1connect"
	"github.com/SingularityCoLabs/aether/internal/buildinfo"
)

type healthCheckerFunc func(context.Context) error

func (function healthCheckerFunc) Ping(ctx context.Context) error {
	return function(ctx)
}

func TestSystemServiceUsesSharedContract(t *testing.T) {
	startedAt := time.Date(2026, time.August, 20, 12, 0, 0, 0, time.UTC)
	server := httptest.NewServer(NewHandler(Options{
		Logger:      slog.New(slog.NewTextHandler(io.Discard, nil)),
		Build:       buildinfo.Info{Version: "v0.0.0-test", Commit: "abc123", Date: "today"},
		Environment: "test",
		StartedAt:   startedAt,
	}))
	defer server.Close()

	client := aetherv1connect.NewSystemServiceClient(http.DefaultClient, server.URL)
	response, err := client.GetSystemInfo(
		context.Background(),
		connect.NewRequest(&aetherv1.GetSystemInfoRequest{}),
	)
	if err != nil {
		t.Fatalf("GetSystemInfo() error = %v", err)
	}
	if response.Msg.GetVersion() != "v0.0.0-test" {
		t.Fatalf("version = %q, want v0.0.0-test", response.Msg.GetVersion())
	}
	if !response.Msg.GetStartedAt().AsTime().Equal(startedAt) {
		t.Fatalf("startedAt = %s, want %s", response.Msg.GetStartedAt().AsTime(), startedAt)
	}
}

func TestReadinessFailsClosedWhenDatabaseIsUnavailable(t *testing.T) {
	handler := NewHandler(Options{
		Database: healthCheckerFunc(func(context.Context) error {
			return errors.New("not connected")
		}),
	})
	request := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusServiceUnavailable)
	}
}

func TestCORSAllowsConfiguredDashboardOnly(t *testing.T) {
	handler := NewHandler(Options{AllowedOrigins: []string{"http://localhost:3000"}})

	allowed := httptest.NewRequest(http.MethodOptions, "/aether.v1.SystemService/GetSystemInfo", nil)
	allowed.Header.Set("Origin", "http://localhost:3000")
	allowedResponse := httptest.NewRecorder()
	handler.ServeHTTP(allowedResponse, allowed)
	if allowedResponse.Code != http.StatusNoContent {
		t.Fatalf("allowed status = %d, want %d", allowedResponse.Code, http.StatusNoContent)
	}

	rejected := httptest.NewRequest(http.MethodOptions, "/aether.v1.SystemService/GetSystemInfo", nil)
	rejected.Header.Set("Origin", "https://attacker.example")
	rejectedResponse := httptest.NewRecorder()
	handler.ServeHTTP(rejectedResponse, rejected)
	if rejectedResponse.Code != http.StatusForbidden {
		t.Fatalf("rejected status = %d, want %d", rejectedResponse.Code, http.StatusForbidden)
	}
}
