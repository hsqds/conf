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

const (
	delimiter  = "-"
	assignment = "="
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

// GetPriority
func (s *FlagsSource) GetPriority() int {
	return s.priority
}

// GetServiceConfig
func (s *FlagsSource) GetServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, fmt.Errorf("could not get config for %q service", serviceName)
	}
	return cfg, nil
}

// Load
func (s *FlagsSource) Load(ctx context.Context, services []string) error {
	var err error
	s.data, err = s.parse(services, os.Args)
	if err != nil {
		return fmt.Errorf("could not parse flags: %w", err)
	}
	return nil
}

// parse gets config parameters from args
func (s *FlagsSource) parse(services []string, args []string) (map[string]conf.Config, error) {
	s.logger.Debug().Interface("services", services).Interface("args", args).Send()
	configs := make(map[string]conf.Config)

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
			splitted := strings.SplitN(keyVal, assignment, 2)

			s.logger.Debug().Interface("splitted", splitted).Send()

			if len(splitted) < 2 {
				continue
			}

			svcConfig[splitted[0]] = splitted[1]
		}
		configs[svc] = svcConfig
	}

	return configs, nil
}
