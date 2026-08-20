// Package security owns authorization vocabulary shared by tools and APIs.
package security

// Permission is a stable, namespaced capability checked by policy.
type Permission string

const (
	PermissionResourcesRead Permission = "resources.read"
	PermissionPlansCreate   Permission = "plans.create"
	PermissionPlansApprove  Permission = "plans.approve"
	PermissionOperationsRun Permission = "operations.run"
)
