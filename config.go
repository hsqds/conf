package conf

import (
	"fmt"
	"strings"
	"text/template"
)

type Getter interface {
	// Get config value by key
	Get(key, defaultValue string) string
}

type Formatter interface {
	// Fmt renders `pattern` filling it with config values
	Fmt(pattern string) (string, error)
}

// Config
type Config interface {
	Getter
	Formatter
}

// MapConfig represents map based config
// Implements Config interface
type MapConfig map[string]string

// Get returns value by key or defaultValue
func (m MapConfig) Get(key, defaultValue string) string {
	val, ok := m[key]
	if !ok {
		return defaultValue
	}

	return val
}

// Fmt renders `pattern` filling it with config values
func (m MapConfig) Fmt(pattern string) (string, error) {
	t, err := template.New("pattern").Parse(pattern)
	if err != nil {
		return "", fmt.Errorf("could not parse pattern: %w", err)
	}

	b := &strings.Builder{}
	err = t.Execute(b, m)
	if err != nil {
		return "", fmt.Errorf("could not render template: %w", err)
	}

	return b.String(), nil
}
