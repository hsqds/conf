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
type SyncedSourcesStorage struct {
	mtx     sync.RWMutex
	sources map[string]Source
}

// newSyncedSources
func NewSyncedSourcesStorage() *SyncedSourcesStorage {
	return &SyncedSourcesStorage{
		sources: make(map[string]Source),
	}
}

// Set
func (s *SyncedSourcesStorage) Set(key string, src Source) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.sources[key]; ok {
		return errors.New("source is not unique")
	}

	s.sources[key] = src

	return nil
}
