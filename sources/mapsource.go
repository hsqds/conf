package sources

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hsqds/conf"
)

// MapSource represents.
type MapSource struct {
	id       string
	data     map[string]conf.Config
	priority int
}

// NewMapSource.
func NewMapSource(priority int, data map[string]conf.Config) *MapSource {
	return &MapSource{
		id:       fmt.Sprintf("map-%s", uuid.NewString()),
		data:     data,
		priority: priority,
	}
}

// ID returns unique source id.
func (s *MapSource) ID() string {
	return s.id
}

// Load.
func (s *MapSource) Load(ctx context.Context, services []string) error {
	return nil
}

// Priority.
func (s *MapSource) Priority() int {
	return s.priority
}

// ServiceConfig.
func (s *MapSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, ServiceConfigError{serviceName, s.id}
	}

	return cfg, nil
}

// Close.
func (s *MapSource) Close(ctx context.Context) {}
