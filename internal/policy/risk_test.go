package policy

import "testing"

func TestRequiresApproval(t *testing.T) {
	tests := []struct {
		risk RiskLevel
		want bool
	}{
		{RiskReadOnly, false},
		{RiskLow, false},
		{RiskMedium, true},
		{RiskHigh, true},
		{RiskCritical, true},
	}

	for _, test := range tests {
		t.Run(test.risk.String(), func(t *testing.T) {
			if got := test.risk.RequiresApproval(); got != test.want {
				t.Fatalf("RequiresApproval() = %t, want %t", got, test.want)
			}
		})
	}
}
