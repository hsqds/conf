package conf

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	defaultLoadTimeout = 3000
)

// Loader
type Loader interface {
	Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult
}

// LoadResult represents
type LoadResult struct {
	SourceID string
	Config   Config
	Err      error
	Service  string
}

// Loader represents
type ConfigsLoader struct{}

// load call source loading
func load(ctx context.Context, src Source, services []string, resultsCh chan<- LoadResult, wg *sync.WaitGroup) {
	log.Debug().Msg("start loading routine")
	defer src.Close(ctx)
	defer wg.Done()

	result := LoadResult{
		SourceID: src.ID(),
	}

	defer func(res *LoadResult) {
		log.Debug().Interface("result", *res).Msg("loading routine done")
	}(&result)

	log.Debug().Msg("Start loading from source")
	err := src.Load(ctx, services)
	if err != nil {
		result.Err = err
		resultsCh <- result
		return
	}

	for _, svc := range services {
		log.Debug().Str("service", svc).Msg("getting service config")
		cfg, err := src.GetServiceConfig(svc)
		if err != nil {
			result.Err = err
			result.Service = svc
			resultsCh <- result
			continue
		}

		result.Config = cfg
		result.Service = svc
		resultsCh <- result
	}
}

// Load loads configs from each source in parallel
// !IMPORTANT: context with timeout wanted here!
func (l *ConfigsLoader) Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult {
	var (
		srcCount  = len(sources)
		resultsCh = make(chan LoadResult, srcCount*len(serviceNames))
		results   = make([]LoadResult, 0, srcCount)
		cancel    func()
	)

	// if no deadline is set create context with default timeout value
	_, ok := ctx.Deadline()
	if !ok {
		log.Debug().Msg("context has no deadline. will set default timeout")
		ctx, cancel = context.WithTimeout(ctx, defaultLoadTimeout*time.Millisecond)
		defer cancel()
	}

	wg := sync.WaitGroup{}
	wg.Add(srcCount)
	log.Debug().Int("src count", srcCount).Msg("wait group value")

	for _, src := range sources {
		log.Debug().Str("source id", src.ID()).Msg("loading from source")
		go load(ctx, src, serviceNames, resultsCh, &wg)
	}

	log.Debug().Msg("waiting fo wg")
	wg.Wait()
	log.Debug().Msg("closing results channel")
	close(resultsCh) // after wg.Wait, so it is ok

	for res := range resultsCh {
		log.Debug().Interface("result", res).Msg("moving result from channel to slice")
		results = append(results, res)
	}

	log.Debug().Interface("results list", results).Send()
	return results
}
