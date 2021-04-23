package provider

import (
	"context"
	"errors"
	"sync"
)

// SourceSetter
type SourcesStorage interface {
	Set(serviceName string, src Source) error
	List() []Source
}

// Source
type Source interface {
	// Shoud return unique source identifier persistent for in all source lifetime
	ID() string
	// Ping let insure source is available
	Ping(context.Context) error
	// Load pull config for the list of service
	Load(context.Context, []string) error
	// GetPriority returns source priority
	GetPriority() int
	// GetServiceConfig
	GetServiceConfig(serviceName string) (Config, error)
	// Close closes connections
	Close(context.Context) error
}

// syncedSources represents sources map protected with mutex
type syncedSourceStorage struct {
	mtx     sync.RWMutex
	sources map[string]Source
}

// newSyncedSources
func newSyncedSourceStorage() *syncedSourceStorage {
	return &syncedSourceStorage{
		sources: make(map[string]Source),
	}
}

// Set
func (sm *syncedSourceStorage) Set(key string, src Source) error {
	sm.mtx.Lock()
	defer sm.mtx.Unlock()

	if _, ok := sm.sources[key]; ok {
		return errors.New("source is not unique")
	}

	sm.sources[key] = src

	return nil
}
