package approval

import (
	"testing"
	"time"
)

func TestExpiredApprovalCannotAuthorizeExecution(t *testing.T) {
	now := time.Now()
	subject := Approval{
		State:     StatePending,
		ExpiresAt: now.Add(-time.Second),
	}

	_, err := subject.Decide(StateApproved, "user-1", now)
	if err == nil {
		t.Fatal("Decide() error = nil, want expiration error")
	}
}

func TestApprovalRecordsActorAndTime(t *testing.T) {
	now := time.Now()
	subject := Approval{
		State:     StatePending,
		ExpiresAt: now.Add(time.Hour),
	}

	got, err := subject.Decide(StateApproved, "user-1", now)
	if err != nil {
		t.Fatalf("Decide() error = %v", err)
	}
	if got.DecidedBy != "user-1" || !got.DecidedAt.Equal(now.UTC()) {
		t.Fatalf("decision metadata = %q/%s", got.DecidedBy, got.DecidedAt)
	}
}
