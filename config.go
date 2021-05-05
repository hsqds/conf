package conf

import (
	"fmt"
	"sync"
)

// Config
type Config interface {
	// Get config value by key
	Get(key string, defaultValue string) string
}

// ConfigsStorage
type ConfigsStorage interface {
	Get(serviceName string) (Config, error)
	Set(serviceName string, cfg Config) error
}

// ConfigsStorage represents configs map protected with mutex
type SyncedConfigsStorage struct {
	mtx     sync.RWMutex
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
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	cfg, ok := c.configs[serviceName]
	if !ok {
		return nil, fmt.Errorf("no config for service %q", serviceName)
	}

	return cfg, nil
}
