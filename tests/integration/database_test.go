package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/SingularityCoLabs/aether/internal/database"
)

func TestEmbeddedMigrations(t *testing.T) {
	databaseURL := os.Getenv("AETHER_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("AETHER_TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := database.Migrate(ctx, databaseURL); err != nil {
		t.Fatalf("Migrate() error = %v", err)
	}
	connection, err := database.Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer connection.Close()

	var phase int
	err = connection.Pool().QueryRow(
		ctx,
		"SELECT (value->>'phase')::int FROM system_metadata WHERE key = 'schema'",
	).Scan(&phase)
	if err != nil {
		t.Fatalf("query system metadata: %v", err)
	}
	if phase != 0 {
		t.Fatalf("phase = %d, want 0", phase)
	}
}
