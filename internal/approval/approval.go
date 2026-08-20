// Package approval defines explicit authorization for risk-bearing plans.
package approval

import (
	"errors"
	"time"

	"github.com/SingularityCoLabs/aether/internal/policy"
)

type State string

const (
	StatePending  State = "PENDING"
	StateApproved State = "APPROVED"
	StateRejected State = "REJECTED"
	StateExpired  State = "EXPIRED"
)

// Approval contains authorization metadata only. It does not execute a plan.
type Approval struct {
	ID          string
	PlanID      string
	ProjectID   string
	Risk        policy.RiskLevel
	State       State
	RequestedBy string
	DecidedBy   string
	RequestedAt time.Time
	DecidedAt   time.Time
	ExpiresAt   time.Time
}

// Decide applies one terminal decision to a pending, unexpired approval.
func (a Approval) Decide(state State, actor string, now time.Time) (Approval, error) {
	if a.State != StatePending {
		return Approval{}, errors.New("only a pending approval can be decided")
	}
	if state != StateApproved && state != StateRejected {
		return Approval{}, errors.New("decision must be approved or rejected")
	}
	if actor == "" {
		return Approval{}, errors.New("decision actor is required")
	}
	if !a.ExpiresAt.IsZero() && !now.Before(a.ExpiresAt) {
		return Approval{}, errors.New("approval has expired")
	}
	a.State = state
	a.DecidedBy = actor
	a.DecidedAt = now.UTC()
	return a, nil
}
