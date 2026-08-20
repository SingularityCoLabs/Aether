package provider

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Registry makes provider selection explicit and rejects ambiguous names.
type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
}

func NewRegistry(providers ...Provider) (*Registry, error) {
	registry := &Registry{providers: make(map[string]Provider, len(providers))}
	for _, candidate := range providers {
		if err := registry.Register(candidate); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (r *Registry) Register(candidate Provider) error {
	if candidate == nil {
		return errors.New("provider is nil")
	}
	name := candidate.Name()
	if name == "" {
		return errors.New("provider name is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.providers[name]; exists {
		return fmt.Errorf("provider %q is already registered", name)
	}
	r.providers[name] = candidate
	return nil
}

func (r *Registry) Get(name string) (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	candidate, ok := r.providers[name]
	return candidate, ok
}

func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.providers))
	for name := range r.providers {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
