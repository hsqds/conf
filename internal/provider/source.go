package provider

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// SourceSetter
type SourcesStorage interface {
	Append(src Source) error
	Get(sourceID string) (Source, error)
	List() []Source
}

// Source
type Source interface {
	// Shoud return unique source identifier persistent for in all source lifetime
	ID() string
	// Ping let insure source is available
	Ping(context.Context) error
	// Load pull config for the list of service
	Load(ctx context.Context, serviceNames []string) error
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
func (s *SyncedSourcesStorage) Append(src Source) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	srcID := src.ID()

	if _, ok := s.sources[srcID]; ok {
		return errors.New("source is not unique")
	}

	s.sources[srcID] = src

	return nil
}

// List returns sources as a slice
func (s *SyncedSourcesStorage) List() []Source {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	lst := make([]Source, 0, len(s.sources))
	for _, src := range s.sources {
		lst = append(lst, src)
	}

	return lst
}

// Get
func (s *SyncedSourcesStorage) Get(sourceID string) (Source, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	src, ok := s.sources[sourceID]
	if !ok {
		return nil, fmt.Errorf("SyncedSourceStorage: no source with id %q", sourceID)
	}

	return src, nil
}
