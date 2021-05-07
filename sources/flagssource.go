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

// FlagsSource represents
type FlagsSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	prefix string

	logger zerolog.Logger
}

// NewFlagsSource
func NewFlagsSource(priority int, prefix string, logger *zerolog.Logger) *FlagsSource {
	*logger = logger.With().Str("component", "source.flags").Logger()
	return &FlagsSource{
		id:       uuid.NewString(),
		data:     make(map[string]conf.Config),
		priority: priority,
		prefix:   prefix,
		logger:   *logger,
	}
}

// funcname
func (s *FlagsSource) Close(ctx context.Context) error {
	return nil
}

// ID
func (s *FlagsSource) ID() string {
	return s.id
}

// Priority
func (s *FlagsSource) Priority() int {
	return s.priority
}

// ServiceConfig
func (s *FlagsSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, fmt.Errorf("could not get config for %q service", serviceName)
	}
	return cfg, nil
}

// Load
func (s *FlagsSource) Load(ctx context.Context, services []string) error {
	s.data = s.parse(services, os.Args)
	return nil
}

// parse gets config parameters from args
func (s *FlagsSource) parse(services, args []string) map[string]conf.Config {
	s.logger.Debug().Interface("services", services).Interface("args", args).Send()
	configs := make(map[string]conf.Config)

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
		for _, arg := range args[1:] {
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

			svcConfig[splitted[0]] = splitted[1]
		}
		configs[svc] = svcConfig
	}

	return configs
}
