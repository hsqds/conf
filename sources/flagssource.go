package sources

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/hsqds/conf"
)

// FlagsSource represents.
type FlagsSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	prefix string

	args []string
}

// NewFlagsSource initializes FlagsSource using standard `os.Args`.
func NewFlagsSource(priority int, prefix string) *FlagsSource {
	s := initFlagsSource(priority, prefix, os.Args)

	return &s
}

// initFlagsSource
func initFlagsSource(priority int, prefix string, args []string) FlagsSource {
	return FlagsSource{
		id:       fmt.Sprintf("flags-%s", uuid.NewString()),
		data:     make(map[string]conf.Config),
		priority: priority,
		prefix:   prefix,
		args:     args,
	}
}

// Close.
func (s *FlagsSource) Close(ctx context.Context) {}

// ID.
func (s *FlagsSource) ID() string {
	return s.id
}

// Priority.
func (s *FlagsSource) Priority() int {
	return s.priority
}

// ServiceConfig.
func (s *FlagsSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, ServiceConfigError{serviceName, s.id}
	}

	return cfg, nil
}

// Load loads configuration data from flags passed at args.
func (s *FlagsSource) Load(ctx context.Context, services []string) (err error) {
	const (
		delimiter  = "-"
		assignment = "="
	)

	for _, svc := range uniqueStrings(services) {
		s.data[svc] = conf.NewMapConfig(sieveServiceConfig(svc, s.prefix, delimiter, assignment, s.args))
	}

	return err
}
