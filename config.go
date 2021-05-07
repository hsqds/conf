package conf

import (
	"fmt"
	"strings"
	"sync"
	"text/template"
)

type Getter interface {
	// Get config value by key
	Get(key, defaultValue string) (string, bool)
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
type MapConfig map[string]string

// Get returns value by key or defaultValue
func (m MapConfig) Get(key, defaultValue string) (string, bool) {
	val, ok := m[key]
	if !ok {
		return defaultValue, false
	}

	return val, true
}

// Fmt renders `pattern` filling it with config values
func (m MapConfig) Fmt(pattern string) (string, error) {
	p, err := template.New("pattern").Parse(pattern)
	if err != nil {
		return "", fmt.Errorf("could not parse pattern: %w", err)
	}

	b := &strings.Builder{}
	err = p.Execute(b, m)
	if err != nil {
		return "", fmt.Errorf("could not render template: %w", err)
	}

	return b.String(), nil
}

// ConfigsStorage
type ConfigsStorage interface {
	Get(serviceName string) (Config, error)
	Set(serviceName string, cfg Config) error
}

// ConfigsStorage represents configs map protected with mutex
type SyncedConfigsStorage struct {
	mtx     sync.Mutex
	configs map[string]Config
}

func NewSyncedConfigsStorage() *SyncedConfigsStorage {
	return &SyncedConfigsStorage{
		configs: make(map[string]Config),
	}
}

// Set
func (c *SyncedConfigsStorage) Set(serviceName string, cfg Config) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.configs[serviceName] = cfg

	return nil
}

// Get receives configs by service name
func (c *SyncedConfigsStorage) Get(serviceName string) (Config, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cfg, ok := c.configs[serviceName]
	if !ok {
		return nil, fmt.Errorf("no config for service %q", serviceName)
	}

	return cfg, nil
}
