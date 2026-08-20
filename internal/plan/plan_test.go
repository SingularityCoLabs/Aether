package plan

import (
	"encoding/json"
	"testing"

	"github.com/SingularityCoLabs/aether/internal/policy"
)

func TestValidateRejectsUnderstatedRisk(t *testing.T) {
	subject := Plan{
		Goal: "delete application",
		Risk: policy.RiskLow,
		Steps: []Step{{
			ID:           "delete",
			Action:       "applications.delete",
			ResourceKind: "Application",
			Provider:     "docker",
			Input:        json.RawMessage(`{}`),
			Risk:         policy.RiskHigh,
		}},
	}

	if err := subject.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want understated risk error")
	}
}

func TestValidateAcceptsTypedPlan(t *testing.T) {
	subject := Plan{
		Goal: "inspect application",
		Risk: policy.RiskReadOnly,
		Steps: []Step{{
			ID:           "inspect",
			Action:       "applications.inspect",
			ResourceKind: "Application",
			Provider:     "docker",
			Input:        json.RawMessage(`{"name":"api"}`),
			Risk:         policy.RiskReadOnly,
		}},
	}

	if err := subject.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}
