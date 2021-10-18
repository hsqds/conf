package conf

import (
	"sync"
)

// ConfigsStorage.
type ConfigsStorage interface {
	ByServiceName(serviceName string) (Config, error)
	Set(serviceName string, cfg Config) error
	Has(serviceName string) bool
}

// ConfigsStorage represents configs map protected with mutex.
type SyncedConfigsStorage struct {
	mtx     sync.Mutex
	configs map[string]Config
}

func NewSyncedConfigsStorage() *SyncedConfigsStorage {
	return &SyncedConfigsStorage{
		configs: make(map[string]Config),
	}
}

// Set.
func (c *SyncedConfigsStorage) Set(serviceName string, cfg Config) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.configs[serviceName] = cfg

	return nil
}

// Has checks service config exist
func (c *SyncedConfigsStorage) Has(serviceName string) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.configs[serviceName]

	return ok
}

// Get receives configs by service name.
func (c *SyncedConfigsStorage) ByServiceName(serviceName string) (Config, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cfg, ok := c.configs[serviceName]
	if !ok {
		return nil, ErrServiceConfigNotFound{serviceName}
	}

	return cfg, nil
}
