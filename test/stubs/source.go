package stubs

import (
	"context"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hsqds/conf"
)

const maxPriority = 20

// TestSource represents basic Source interface
// implementation for testing.
type TestSource struct {
	p  int
	id string
}

// NewTestSource.
func NewTestSource() TestSource {
	return TestSource{
		p:  rand.Intn(maxPriority),
		id: uuid.NewString(),
	}
}

// ID.
func (s *TestSource) ID() string {
	return s.id
}

// Ping.
func (s *TestSource) Ping(ctx context.Context) error {
	return nil
}

// Load.
func (s *TestSource) Load(ctx context.Context, serviceNames []string) error {
	return nil
}

// funcname.
func (s *TestSource) Priority() int {
	return s.p
}

// ServiceConfig.
func (s *TestSource) ServiceConfig(serviceName string) (conf.Config, error) {
	return nil, nil
}

// Close.
func (s *TestSource) Close(context.Context) error {
	return nil
}
