package sources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/hsqds/conf"
)

// SrcReader
type SrcReader func(io.Reader) ([]byte, error)

// JSONUnmarshaler
type JSONUnmarshaler func([]byte, interface{}) error

// JSONFileStorage represents.
type JSONSource struct {
	id       string
	data     map[string]conf.Config
	priority int

	src       io.ReadCloser
	srcRead   SrcReader
	unmarshal JSONUnmarshaler
}

// NewJSONSource
func NewJSONSource(priority int, src io.ReadCloser) *JSONSource {
	s := JSONSource{
		id:        fmt.Sprintf("json-%s", uuid.NewString()),
		data:      make(map[string]conf.Config),
		priority:  priority,
		src:       src,
		srcRead:   io.ReadAll,
		unmarshal: json.Unmarshal,
	}

	return &s
}

// SetReader
func (s *JSONSource) SetReader(r SrcReader) {
	s.srcRead = r
}

// SetJSONUnmarshaler
func (s *JSONSource) SetJSONUnmarshaler(u JSONUnmarshaler) {
	s.unmarshal = u
}

// ID
func (s *JSONSource) ID() string {
	return s.id
}

// Load
func (s *JSONSource) Load(ctx context.Context, services []string) error {
	raw, err := s.srcRead(s.src)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	tmp := make(map[string]map[string]string)

	err = s.unmarshal(raw, &tmp)
	if err != nil {
		return fmt.Errorf("could not unmarshal json: %w", err)
	}

	for k := range tmp {
		s.data[k] = conf.NewMapConfig(tmp[k])
	}

	return nil
}

// Close
func (s *JSONSource) Close(ctx context.Context) {
	if err := s.src.Close(); err != nil {
		// TODO: log
		_ = err
	}
}

// Priority
func (s *JSONSource) Priority() int {
	return s.priority
}

// ServiceConfig
func (s *JSONSource) ServiceConfig(serviceName string) (conf.Config, error) {
	cfg, ok := s.data[serviceName]
	if !ok {
		return nil, ServiceConfigError{serviceName, s.id}
	}

	return cfg, nil
}
