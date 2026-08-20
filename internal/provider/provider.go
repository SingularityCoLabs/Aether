// Package provider defines the adapter boundary for infrastructure substrates.
package provider

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SingularityCoLabs/aether/internal/plan"
	"github.com/SingularityCoLabs/aether/pkg/resource"
)

// Capability is a typed action implemented by a provider.
type Capability string

// ResourceState is provider-observed actual state.
type ResourceState struct {
	Exists             bool
	ObservedGeneration int64
	Status             json.RawMessage
	ObservedAt         time.Time
}

// Result is the typed outcome of applying one plan step.
type Result struct {
	ExternalID string
	Output     json.RawMessage
}

// Provider is implemented by Docker first and future substrates later.
type Provider interface {
	Name() string
	Capabilities(context.Context) ([]Capability, error)
	Plan(context.Context, resource.Resource, ResourceState) (plan.Plan, error)
	Apply(context.Context, plan.Step) (Result, error)
	Observe(context.Context, resource.Resource) (ResourceState, error)
	Delete(context.Context, resource.Resource) error
}
