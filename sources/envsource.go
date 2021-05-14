package sources

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/uuid"

	"github.com/hsqds/conf"
)

// EnvSource represents.
type EnvSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	envs []string
}

// NewEnvSource.
func NewEnvSource(priority int, envs []string) *EnvSource {
	return &EnvSource{
		id:       fmt.Sprintf("env-%s", uuid.NewString()),
		data:     make(map[string]conf.Config),
		priority: priority,
		envs:     envs,
	}
}

// Close.
func (s *EnvSource) Close(ctx context.Context) {}

// ID.
func (s *EnvSource) ID() string {
	return s.id
}

// Priority.
func (s *EnvSource) Priority() int {
	return s.priority
}

// ServiceConfig.
func (s *EnvSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, ServiceConfigError{serviceName, s.id}
	}

	return cfg, nil
}

// Load.
func (s *EnvSource) Load(ctx context.Context, services []string) error {
	const (
		delimiter  = "_"
		assignment = "="
		matchesNum = 4
		reKeyIndex = 2
		reValIndex = 3
	)

	services = uniqueStrings(services)

	for _, svc := range services {
		pattern := fmt.Sprintf(`((?i)%s%s)([\w]+)%s([\w]+)`, svc, delimiter, assignment)
		svcRe := regexp.MustCompile(pattern)

		svcConfig := conf.MapConfig{}

		for _, kv := range s.envs {
			matches := svcRe.FindStringSubmatch(kv)
			if len(matches) < matchesNum {
				continue
			}

			key := toCamelCase(matches[reKeyIndex], delimiter)
			svcConfig.Set(key, matches[reValIndex])
		}

		s.data[svc] = svcConfig
	}

	return nil
}
