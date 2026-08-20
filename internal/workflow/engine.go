// Package workflow defines orchestration without committing Phase 0 to a
// distributed workflow system.
package workflow

import (
	"context"
	"encoding/json"
)

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusRunning   Status = "RUNNING"
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
	StatusCanceled  Status = "CANCELED"
)

type Execution struct {
	ID     string
	Status Status
	Output json.RawMessage
}

// Engine can be implemented in-process first and by Temporal when real
// long-running workflow requirements arrive.
type Engine interface {
	Start(ctx context.Context, workflow string, input json.RawMessage) (Execution, error)
	Cancel(ctx context.Context, executionID string) error
	Status(ctx context.Context, executionID string) (Execution, error)
}
