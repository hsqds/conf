package sources

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/hsqds/conf"
)

// FlagsSource represents.
type FlagsSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	prefix string

	logger zerolog.Logger
}

// NewFlagsSource.
func NewFlagsSource(priority int, prefix string, logger *zerolog.Logger) *FlagsSource {
	*logger = logger.With().Str("component", "source.flags").Logger()

	return &FlagsSource{
		id:       fmt.Sprintf("flags-%s", uuid.NewString()),
		data:     make(map[string]conf.Config),
		priority: priority,
		prefix:   prefix,
		logger:   *logger,
	}
}

// Close.
func (s *FlagsSource) Close(ctx context.Context) {
}

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
	s.logger.Debug().Interface("services", services).Interface("args", os.Args).Send()

	const (
		delimiter     = "-"
		assignment    = "="
		splittedCount = 2
	)

	var svcPrefix string
	for _, svc := range services {
		svcPrefix = fmt.Sprintf("%s%s%s", s.prefix, svc, delimiter)
		s.logger.Debug().Str("service prefix", svcPrefix).Send()

		svcConfig := conf.MapConfig{}

		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, svcPrefix) {
				continue
			}

			keyVal := strings.Replace(arg, svcPrefix, "", 1)
			splitted := strings.SplitN(keyVal, assignment, splittedCount)

			s.logger.Debug().Interface("splitted", splitted).Send()

			// yet ignoring bool flags
			if len(splitted) < splittedCount {
				continue
			}

			key := toCamelCase(splitted[0], delimiter)
			svcConfig[key] = splitted[1]
		}

		s.data[svc] = svcConfig
	}

	return err
}
