# Agent boundary

The agent is intentionally not executable in Phase 0.

When introduced, this module will contain provider-neutral model adapters,
grounded context assembly, a typed tool registry, plan generation, and
guardrails. Read-only tools arrive first. A tool may never bypass API
authorization, policy, approval, operation, or audit boundaries.
