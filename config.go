package conf

import (
	"fmt"
	"strings"
	"sync"
	"text/template"
)

// Config
type Config interface {
	// Get config value by key
	Get(key, defaultValue string) string
	// Set sets key value
	Set(key, value string)
	// Fmt renders `pattern` filling it with config values
	Fmt(pattern string) (string, error)
}

// MapConfig represents map based config
// Implements Config interface.
type MapConfig struct {
	data map[string]string
	mtx  sync.Mutex
}

// NewMapConfig
func NewMapConfig(d map[string]string) *MapConfig {
	c := MapConfig{
		data: make(map[string]string),
	}

	if d == nil {
		return &c
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	for k := range d {
		c.data[k] = d[k]
	}

	return &c
}

// Get returns value by key or defaultValue.
func (m *MapConfig) Get(key, defaultValue string) string {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	val, ok := m.data[key]
	if !ok {
		return defaultValue
	}

	return val
}

// Set sets key value
func (m *MapConfig) Set(key, value string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.data[key] = value
}

// Fmt renders `pattern` filling it with config values.
func (m *MapConfig) Fmt(pattern string) (string, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	t, err := template.New("pattern").Parse(pattern)
	if err != nil {
		return "", fmt.Errorf("could not parse pattern: %w", err)
	}

	b := &strings.Builder{}

	err = t.Execute(b, m.data)
	if err != nil {
		return "", fmt.Errorf("could not render template: %w", err)
	}

	return b.String(), nil
}
