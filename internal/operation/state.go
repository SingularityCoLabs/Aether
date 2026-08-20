// Package operation defines durable infrastructure-operation semantics.
package operation

import "fmt"

// State is the lifecycle state of an infrastructure operation.
type State string

const (
	StatePending         State = "PENDING"
	StatePlanning        State = "PLANNING"
	StateWaitingApproval State = "WAITING_APPROVAL"
	StateRunning         State = "RUNNING"
	StateVerifying       State = "VERIFYING"
	StateSucceeded       State = "SUCCEEDED"
	StateFailed          State = "FAILED"
	StateRollingBack     State = "ROLLING_BACK"
	StateRolledBack      State = "ROLLED_BACK"
	StateCanceled        State = "CANCELED"
)

var transitions = map[State]map[State]struct{}{
	StatePending: {
		StatePlanning: {},
		StateCanceled: {},
		StateFailed:   {},
	},
	StatePlanning: {
		StateWaitingApproval: {},
		StateRunning:         {},
		StateCanceled:        {},
		StateFailed:          {},
	},
	StateWaitingApproval: {
		StateRunning:  {},
		StateCanceled: {},
		StateFailed:   {},
	},
	StateRunning: {
		StateVerifying: {},
		StateCanceled:  {},
		StateFailed:    {},
	},
	StateVerifying: {
		StateSucceeded: {},
		StateFailed:    {},
	},
	StateFailed: {
		StateRollingBack: {},
	},
	StateRollingBack: {
		StateRolledBack: {},
		StateFailed:     {},
	},
}

// CanTransition reports whether the state machine permits the transition.
func CanTransition(from, to State) bool {
	_, ok := transitions[from][to]
	return ok
}

// Transition returns the next state or an error without mutating an operation.
func Transition(from, to State) (State, error) {
	if !CanTransition(from, to) {
		return from, fmt.Errorf("operation cannot transition from %s to %s", from, to)
	}
	return to, nil
}

// Terminal reports whether no further successful transition is possible.
func (s State) Terminal() bool {
	return s == StateSucceeded || s == StateRolledBack || s == StateCanceled
}
