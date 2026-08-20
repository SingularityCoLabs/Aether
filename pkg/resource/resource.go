// Package resource defines Aether's provider-independent desired-state model.
package resource

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	kindPattern = regexp.MustCompile(`^[A-Z][A-Za-z0-9]{0,62}$`)
	namePattern = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9.-]{0,61}[a-z0-9])?$`)
)

// Resource is the common vocabulary shared by applications, databases, routes,
// nodes, and future infrastructure kinds.
type Resource struct {
	ID                 string          `json:"id,omitempty"`
	Kind               string          `json:"kind"`
	Name               string          `json:"name"`
	ProjectID          string          `json:"projectId"`
	Spec               json.RawMessage `json:"spec"`
	Status             json.RawMessage `json:"status"`
	Generation         int64           `json:"generation"`
	ObservedGeneration int64           `json:"observedGeneration"`
	Provider           string          `json:"provider"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
}

// New returns a validated resource with an empty observed status.
func New(projectID, kind, name, provider string, spec json.RawMessage, now time.Time) (Resource, error) {
	r := Resource{
		ProjectID:  projectID,
		Kind:       kind,
		Name:       name,
		Spec:       cloneJSON(spec),
		Status:     json.RawMessage("{}"),
		Generation: 1,
		Provider:   provider,
		CreatedAt:  now.UTC(),
		UpdatedAt:  now.UTC(),
	}
	if err := r.Validate(); err != nil {
		return Resource{}, err
	}
	return r, nil
}

// Validate checks provider-independent invariants. Kind-specific validation
// belongs to a resource definition, not a provider adapter.
func (r Resource) Validate() error {
	switch {
	case r.ProjectID == "":
		return errors.New("project ID is required")
	case !kindPattern.MatchString(r.Kind):
		return errors.New("kind must be PascalCase and at most 63 characters")
	case !namePattern.MatchString(r.Name):
		return errors.New("name must be a lowercase DNS-style name and at most 63 characters")
	case r.Provider == "":
		return errors.New("provider is required")
	case r.Generation < 1:
		return errors.New("generation must be at least 1")
	case r.ObservedGeneration < 0:
		return errors.New("observed generation cannot be negative")
	case r.ObservedGeneration > r.Generation:
		return errors.New("observed generation cannot exceed generation")
	}
	if err := validateJSONObject("spec", r.Spec); err != nil {
		return err
	}
	if err := validateJSONObject("status", r.Status); err != nil {
		return err
	}
	if !r.CreatedAt.IsZero() && !r.UpdatedAt.IsZero() && r.UpdatedAt.Before(r.CreatedAt) {
		return errors.New("updated time cannot be before created time")
	}
	return nil
}

// WithSpec returns an updated copy. Semantically identical JSON does not
// advance the generation.
func (r Resource) WithSpec(spec json.RawMessage, now time.Time) (Resource, error) {
	if err := validateJSONObject("spec", spec); err != nil {
		return Resource{}, err
	}
	equal, err := jsonEqual(r.Spec, spec)
	if err != nil {
		return Resource{}, fmt.Errorf("compare spec: %w", err)
	}
	if equal {
		return r, nil
	}
	r.Spec = cloneJSON(spec)
	r.Generation++
	r.UpdatedAt = now.UTC()
	if err := r.Validate(); err != nil {
		return Resource{}, err
	}
	return r, nil
}

// Drifted reports whether the provider has observed the latest desired state.
func (r Resource) Drifted() bool {
	return r.ObservedGeneration < r.Generation
}

func validateJSONObject(field string, value json.RawMessage) error {
	if len(bytes.TrimSpace(value)) == 0 {
		return fmt.Errorf("%s must be a JSON object", field)
	}
	var object map[string]any
	if err := json.Unmarshal(value, &object); err != nil {
		return fmt.Errorf("%s must be a JSON object: %w", field, err)
	}
	if object == nil {
		return fmt.Errorf("%s must be a JSON object", field)
	}
	return nil
}

func jsonEqual(left, right json.RawMessage) (bool, error) {
	var leftValue, rightValue map[string]any
	if err := json.Unmarshal(left, &leftValue); err != nil {
		return false, err
	}
	if err := json.Unmarshal(right, &rightValue); err != nil {
		return false, err
	}
	leftCanonical, err := json.Marshal(leftValue)
	if err != nil {
		return false, err
	}
	rightCanonical, err := json.Marshal(rightValue)
	if err != nil {
		return false, err
	}
	return bytes.Equal(leftCanonical, rightCanonical), nil
}

func cloneJSON(value json.RawMessage) json.RawMessage {
	return append(json.RawMessage(nil), value...)
}
