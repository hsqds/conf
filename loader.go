package conf

import (
	"context"
	"strings"
	"sync"
)

// LoadResult represents.
type LoadResult struct {
	SourceID string
	Config   Config
	Err      error
	Service  string
	Priority int
}

// Loader.
type Loader interface {
	Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult
}

// Loader represents.
type ConfigsLoader int

// load call source loading.
func (cl *ConfigsLoader) load(ctx context.Context, src Source, services []string, resultsCh chan<- LoadResult,
	wg *sync.WaitGroup) {
	defer src.Close(ctx)

	defer wg.Done()

	result := LoadResult{
		SourceID: src.ID(),
	}

	err := src.Load(ctx, services)
	if err != nil {
		result.Err = LoadError{
			Service:  strings.Join(services, ", "),
			SourceID: src.ID(),
			Err:      err,
		}
		resultsCh <- result

		return
	}

	for _, svc := range services {
		cfg, err := src.ServiceConfig(svc)
		if err != nil {
			result.Err = LoadError{
				Service:  svc,
				SourceID: src.ID(),
				Err:      err,
			}
			result.Service = svc
			resultsCh <- result

			continue
		}

		result.Config = cfg
		result.Service = svc
		result.Priority = src.Priority()
		resultsCh <- result
	}
}

// Load loads configs from each source in parallel
func (cl *ConfigsLoader) Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult {
	var (
		srcCount  = len(sources)
		resultsCh = make(chan LoadResult, srcCount*len(serviceNames))
		results   = make([]LoadResult, 0, srcCount)
	)

	wg := sync.WaitGroup{}
	wg.Add(srcCount)

	for _, src := range sources {
		go cl.load(ctx, src, serviceNames, resultsCh, &wg)
	}

	wg.Wait()
	close(resultsCh) // after wg.Wait, so it is ok

	for res := range resultsCh {
		results = append(results, res)
	}

	return results
}
