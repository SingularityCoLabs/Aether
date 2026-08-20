// Package audit defines append-only security and operation evidence.
package audit

import (
	"context"
	"encoding/json"
	"time"
)

// Event records who intended what, in which scope, and with which outcome.
// Metadata must contain references or redacted values, never secret material.
type Event struct {
	ID         string
	ActorID    string
	Action     string
	ProjectID  string
	ResourceID string
	Outcome    string
	Metadata   json.RawMessage
	OccurredAt time.Time
}

// Recorder is implemented by durable audit storage in a later capability
// phase. Mutating application services will depend on this interface.
type Recorder interface {
	Record(context.Context, Event) error
}
