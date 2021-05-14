package conf

import (
	"context"
	"sync"
)

// SourceSetter.
type SourcesStorage interface {
	Append(src Source) error
	Get(sourceID string) (Source, error)
	List() []Source
}

// Source.
type Source interface {
	// Shoud return unique source identifier persistent for in all source lifetime
	ID() string
	// Load pull config for the list of service
	Load(ctx context.Context, serviceNames []string) error
	// Priority returns source priority
	Priority() int
	// ServiceConfig
	ServiceConfig(serviceName string) (Config, error)
	// Close closes connections
	Close(context.Context)
}

// syncedSources represents sources map protected with mutex.
type SyncedSourcesStorage struct {
	mtx     sync.Mutex
	sources map[string]Source
}

// newSyncedSources.
func NewSyncedSourcesStorage() *SyncedSourcesStorage {
	return &SyncedSourcesStorage{
		sources: make(map[string]Source),
	}
}

// Set.
func (s *SyncedSourcesStorage) Append(src Source) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	srcID := src.ID()

	if _, ok := s.sources[srcID]; ok {
		return SourceUniquenessError{srcID}
	}

	s.sources[srcID] = src

	return nil
}

// List returns sources as a slice.
func (s *SyncedSourcesStorage) List() []Source {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	lst := make([]Source, 0, len(s.sources))
	for _, src := range s.sources {
		lst = append(lst, src)
	}

	return lst
}

// Get.
func (s *SyncedSourcesStorage) Get(sourceID string) (Source, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	src, ok := s.sources[sourceID]
	if !ok {
		return nil, SourceStorageError{sourceID}
	}

	return src, nil
}
