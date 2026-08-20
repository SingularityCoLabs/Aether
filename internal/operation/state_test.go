package operation

import "testing"

func TestHappyPath(t *testing.T) {
	path := []State{
		StatePending,
		StatePlanning,
		StateWaitingApproval,
		StateRunning,
		StateVerifying,
		StateSucceeded,
	}

	for i := 0; i < len(path)-1; i++ {
		got, err := Transition(path[i], path[i+1])
		if err != nil {
			t.Fatalf("Transition(%s, %s) error = %v", path[i], path[i+1], err)
		}
		if got != path[i+1] {
			t.Fatalf("Transition(%s, %s) = %s", path[i], path[i+1], got)
		}
	}
}

func TestCannotSkipVerification(t *testing.T) {
	got, err := Transition(StateRunning, StateSucceeded)
	if err == nil {
		t.Fatal("Transition() error = nil, want invalid transition error")
	}
	if got != StateRunning {
		t.Fatalf("Transition() state = %s, want unchanged RUNNING", got)
	}
}

func TestTerminalStates(t *testing.T) {
	for _, state := range []State{StateSucceeded, StateRolledBack, StateCanceled} {
		if !state.Terminal() {
			t.Errorf("%s.Terminal() = false, want true", state)
		}
	}
	if StateFailed.Terminal() {
		t.Error("FAILED.Terminal() = true, want false because rollback remains possible")
	}
}
