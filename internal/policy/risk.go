// Package policy owns risk classification and approval policy contracts.
package policy

import "fmt"

// RiskLevel classifies the maximum impact of a proposed action.
type RiskLevel uint8

const (
	RiskReadOnly RiskLevel = iota
	RiskLow
	RiskMedium
	RiskHigh
	RiskCritical
)

func (r RiskLevel) String() string {
	switch r {
	case RiskReadOnly:
		return "read_only"
	case RiskLow:
		return "low"
	case RiskMedium:
		return "medium"
	case RiskHigh:
		return "high"
	case RiskCritical:
		return "critical"
	default:
		return fmt.Sprintf("unknown(%d)", r)
	}
}

// RequiresApproval is deliberately conservative until tenant-aware policy is
// implemented: every medium-or-higher action needs an explicit approval.
func (r RiskLevel) RequiresApproval() bool {
	return r >= RiskMedium
}
