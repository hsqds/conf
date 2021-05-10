package sources

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/hsqds/conf"
	"github.com/rs/zerolog"
)

// EnvSource represents
type EnvSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	logger zerolog.Logger
}

// NewEnvSource
func NewEnvSource(priority int, logger *zerolog.Logger) *EnvSource {
	*logger = logger.With().Str("component", "source.env").Logger()
	return &EnvSource{
		id:       uuid.NewString(),
		data:     make(map[string]conf.Config),
		priority: priority,
		logger:   *logger,
	}
}

// Close
func (s *EnvSource) Close(ctx context.Context) error {
	return nil
}

// ID
func (s *EnvSource) ID() string {
	return s.id
}

// Priority
func (s *EnvSource) Priority() int {
	return s.priority
}

// ServiceConfig
func (s *EnvSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, fmt.Errorf("could not get config for %s service", serviceName)
	}

	return cfg, nil
}

// Load
func (s *EnvSource) Load(ctx context.Context, services []string) error {
	// TODO: check what is faster: regex or iteration over services names
	const (
		delimiter     = "_"
		assignment    = "="
		splittedCount = 2
	)

	envs := os.Environ()

	for _, svc := range services {
		svcPrefix := fmt.Sprintf("%s%s", strings.ToUpper(svc), delimiter)
		svcConfig := conf.MapConfig{}
		for _, kv := range envs {
			if !strings.HasPrefix(kv, svcPrefix) {
				continue
			}
			kv = strings.ReplaceAll(kv, svcPrefix, "")
			s.logger.Debug().Msg(kv)
			splitted := strings.SplitN(kv, assignment, splittedCount)

			if len(splitted) < splittedCount {
				continue
			}

			s.logger.Debug().Interface("splitted", splitted).Send()

			key := toCamelCase(splitted[0], delimiter)
			s.logger.Debug().Str("cc", key).Send()

			svcConfig[key] = splitted[1]
		}

		s.logger.Debug().Str("service", svc).Interface("service config", svcConfig).Send()

		s.data[svc] = svcConfig
	}

	return nil
}
