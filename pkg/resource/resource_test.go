package resource

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewResourceStartsDrifted(t *testing.T) {
	now := time.Date(2026, time.August, 20, 10, 0, 0, 0, time.FixedZone("LKT", 5*60*60+30*60))
	got, err := New("project-1", "Application", "api", "docker", json.RawMessage(`{"image":"nginx:1.29"}`), now)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if got.Generation != 1 || got.ObservedGeneration != 0 {
		t.Fatalf("generations = %d/%d, want 1/0", got.Generation, got.ObservedGeneration)
	}
	if !got.Drifted() {
		t.Fatal("Drifted() = false, want true")
	}
	if got.CreatedAt.Location() != time.UTC {
		t.Fatalf("CreatedAt location = %v, want UTC", got.CreatedAt.Location())
	}
}

func TestWithSpecOnlyAdvancesForSemanticChange(t *testing.T) {
	now := time.Now()
	original, err := New("project-1", "Application", "api", "docker", json.RawMessage(`{"image":"nginx","replicas":1}`), now)
	if err != nil {
		t.Fatal(err)
	}

	unchanged, err := original.WithSpec(json.RawMessage(`{"replicas":1,"image":"nginx"}`), now.Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	if unchanged.Generation != 1 {
		t.Fatalf("unchanged generation = %d, want 1", unchanged.Generation)
	}

	changed, err := original.WithSpec(json.RawMessage(`{"image":"nginx","replicas":2}`), now.Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	if changed.Generation != 2 {
		t.Fatalf("changed generation = %d, want 2", changed.Generation)
	}
}

func TestValidateRejectsNonObjectSpec(t *testing.T) {
	_, err := New("project-1", "Application", "api", "docker", json.RawMessage(`["nginx"]`), time.Now())
	if err == nil {
		t.Fatal("New() error = nil, want object validation error")
	}
}

func TestValidateRejectsObservedGenerationAhead(t *testing.T) {
	resource, err := New("project-1", "Application", "api", "docker", json.RawMessage(`{}`), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	resource.ObservedGeneration = 2
	if err := resource.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want generation validation error")
	}
}
