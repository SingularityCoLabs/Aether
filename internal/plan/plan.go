// Package plan models validated infrastructure intent before execution.
package plan

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SingularityCoLabs/aether/internal/policy"
)

// Step is one typed provider operation.
type Step struct {
	ID           string           `json:"id"`
	Action       string           `json:"action"`
	ResourceID   string           `json:"resourceId,omitempty"`
	ResourceKind string           `json:"resourceKind"`
	Provider     string           `json:"provider"`
	Input        json.RawMessage  `json:"input"`
	Risk         policy.RiskLevel `json:"risk"`
}

// Plan is an immutable proposal until it has passed validation and policy.
type Plan struct {
	Goal  string           `json:"goal"`
	Risk  policy.RiskLevel `json:"risk"`
	Steps []Step           `json:"steps"`
}

// Validate rejects incomplete plans and verifies that the aggregate risk is
// never lower than any individual step.
func (p Plan) Validate() error {
	if p.Goal == "" {
		return errors.New("goal is required")
	}
	if len(p.Steps) == 0 {
		return errors.New("at least one step is required")
	}
	seen := make(map[string]struct{}, len(p.Steps))
	var highest policy.RiskLevel
	for i, step := range p.Steps {
		if step.ID == "" {
			return fmt.Errorf("step %d: ID is required", i)
		}
		if _, exists := seen[step.ID]; exists {
			return fmt.Errorf("step %d: duplicate ID %q", i, step.ID)
		}
		seen[step.ID] = struct{}{}
		if step.Action == "" || step.ResourceKind == "" || step.Provider == "" {
			return fmt.Errorf("step %q: action, resource kind, and provider are required", step.ID)
		}
		if len(step.Input) == 0 || !json.Valid(step.Input) {
			return fmt.Errorf("step %q: input must be valid JSON", step.ID)
		}
		if step.Risk > highest {
			highest = step.Risk
		}
	}
	if p.Risk < highest {
		return fmt.Errorf("plan risk %s is lower than step risk %s", p.Risk, highest)
	}
	return nil
}
