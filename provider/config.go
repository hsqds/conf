package provider

import (
	"fmt"
	"sync"
)

// Config
type Config interface {
	// Get config value by key
	Get(key string, defaultValue string) (string, error)
}

// ConfigsStorage
type ConfigsStorage interface {
	Get(serviceName string) (Config, error)
	Update(serviceName string, cfg Config) error
}

// ConfigsStorage represents configs map protected with mutex
type SyncedConfigsStorage struct {
	mtx     sync.RWMutex
	configs map[string]Config
}

func NewSyncedConfigs() *SyncedConfigsStorage {
	return &SyncedConfigsStorage{
		configs: make(map[string]Config),
	}
}

// Update configsCache
func (c *SyncedConfigsStorage) Update(serviceName string, cfg Config) error {
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
