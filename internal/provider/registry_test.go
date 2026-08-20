package provider

import (
	"context"
	"testing"

	"github.com/SingularityCoLabs/aether/internal/plan"
	"github.com/SingularityCoLabs/aether/pkg/resource"
)

type stubProvider struct{ name string }

func (s stubProvider) Name() string { return s.name }
func (stubProvider) Capabilities(context.Context) ([]Capability, error) {
	return nil, nil
}
func (stubProvider) Plan(context.Context, resource.Resource, ResourceState) (plan.Plan, error) {
	return plan.Plan{}, nil
}
func (stubProvider) Apply(context.Context, plan.Step) (Result, error) {
	return Result{}, nil
}
func (stubProvider) Observe(context.Context, resource.Resource) (ResourceState, error) {
	return ResourceState{}, nil
}
func (stubProvider) Delete(context.Context, resource.Resource) error { return nil }

func TestRegistryRejectsDuplicateNames(t *testing.T) {
	_, err := NewRegistry(stubProvider{name: "docker"}, stubProvider{name: "docker"})
	if err == nil {
		t.Fatal("NewRegistry() error = nil, want duplicate name error")
	}
}

func TestRegistryGetsProvider(t *testing.T) {
	registry, err := NewRegistry(stubProvider{name: "docker"})
	if err != nil {
		t.Fatal(err)
	}
	got, ok := registry.Get("docker")
	if !ok || got.Name() != "docker" {
		t.Fatalf("Get(docker) = %v, %t", got, ok)
	}
}
