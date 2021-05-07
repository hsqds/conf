package sources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/hsqds/conf"
)

// JSONFileStorage represents
type JSONFileSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	filename string
	file     *os.File

	logger zerolog.Logger
}

// NewJSONFileSource
func NewJSONFileSource(priority int, filename string, logger *zerolog.Logger) *JSONFileSource {
	*logger = logger.With().Caller().Str("component", "source.jsonfile").Logger()
	return &JSONFileSource{
		id:       uuid.NewString(),
		data:     make(map[string]conf.Config),
		priority: priority,
		filename: filename,
		logger:   *logger,
	}
}

// ID
func (s *JSONFileSource) ID() string {
	return s.id
}

// Load
func (s *JSONFileSource) Load(ctx context.Context, services []string) error {
	var err error
	s.file, err = os.Open(s.filename)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	raw, err := io.ReadAll(s.file)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	s.logger.Debug().Str("raw", string(raw)).Send()

	tmp := make(map[string]conf.MapConfig)
	err = json.Unmarshal(raw, &tmp)
	if err != nil {
		return fmt.Errorf("could not unmarshal json: %w", err)
	}

	for k, v := range tmp {
		s.data[k] = conf.Config(v)
	}

	s.logger.Debug().Interface("loaded", s.data).Send()

	return nil
}

// Close
func (s *JSONFileSource) Close(ctx context.Context) error {
	err := s.file.Close()
	if err != nil {
		return fmt.Errorf("could not close source file: %w", err)
	}

	return nil
}

// Priority
func (s *JSONFileSource) Priority() int {
	return s.priority
}

// ServiceConfig
func (s *JSONFileSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, fmt.Errorf("could not get config for %q service", serviceName)
	}
	return cfg, nil
}
