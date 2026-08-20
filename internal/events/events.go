// Package events defines an in-process-first event publishing seam.
package events

import (
	"context"
	"encoding/json"
	"time"
)

type Event struct {
	Name        string
	AggregateID string
	Payload     json.RawMessage
	OccurredAt  time.Time
}

// Publisher can be backed in-process first and by NATS only when independent
// processes require durable delivery.
type Publisher interface {
	Publish(context.Context, Event) error
}
