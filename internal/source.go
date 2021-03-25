package internal

import "context"

// Source
type ISource interface {
	GetAll(ctx context.Context) (map[string]string, error)
	Get(ctx context.Context, key string) (string, error)
}

// PrioritizedSource represents
type PrioritizedSource struct {
	priority int
	source   ISource
}

// NewPrioritizedSource
func NewPrioritizedSource(prio int, src ISource) *PrioritizedSource {
	return &PrioritizedSource{
		priority: prio,
		source:   src,
	}
}
