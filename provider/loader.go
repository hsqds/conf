package provider

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
)

// Loader
type Loader interface {
	Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult
}

// LoadResult represents
type LoadResult struct {
	Source Source
	Config Config
	Err    error
}

// Loader represents
type ConfigsLoader struct {
	mtx sync.Mutex
}

// LoadEach loads configs from each source in parallel
func (l *ConfigsLoader) Load(ctx context.Context, sources []Source, serviceNames []string) <-chan LoadResult {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	count := len(sources)
	results := make(chan LoadResult, count)
	defer close(results)

	wg := sync.WaitGroup{}
	wg.Add(count)

	for _, src := range sources {
		log.Debug().Str("source id", src.ID()).Msg("loading from source")
		go func(src Source, wg sync.WaitGroup) {
			defer src.Close(ctx)
			err := src.Load(ctx, serviceNames)
			var result LoadResult
			if err != nil {
				result = LoadResult{Source: src, Err: err}
			}
			results <- result
			wg.Done()
		}(src, wg)
	}

	wg.Wait()

	return results
}
