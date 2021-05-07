package conf

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"
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
type ConfigsLoader struct {
	logger zerolog.Logger
}

// NewConfigsLoader
func NewConfigsLoader(logger *zerolog.Logger) *ConfigsLoader {
	*logger = logger.With().Caller().Str("component", "conf.loader").Logger().Level(zerolog.DebugLevel)

	return &ConfigsLoader{*logger}
}

// load call source loading
func (cl *ConfigsLoader) load(ctx context.Context, src Source, services []string, resultsCh chan<- LoadResult, wg *sync.WaitGroup) {
	cl.logger.Debug().Msg("start loading routine")
	defer src.Close(ctx)
	defer wg.Done()

	result := LoadResult{
		SourceID: src.ID(),
	}

	defer func(res *LoadResult) {
		cl.logger.Debug().Interface("result", *res).Msg("loading routine done")
	}(&result)

	cl.logger.Debug().Msg("Start loading from source")
	err := src.Load(ctx, services)
	if err != nil {
		result.Err = err
		resultsCh <- result
		return
	}

	for _, svc := range services {
		cl.logger.Debug().Str("service", svc).Msg("getting service config")
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
func (cl *ConfigsLoader) Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult {
	var (
		srcCount  = len(sources)
		resultsCh = make(chan LoadResult, srcCount*len(serviceNames))
		results   = make([]LoadResult, 0, srcCount)
		cancel    func()
	)

	// if no deadline is set create context with default timeout value
	_, ok := ctx.Deadline()
	if !ok {
		cl.logger.Debug().Msg("context has no deadline. will set default timeout")
		ctx, cancel = context.WithTimeout(ctx, defaultLoadTimeout*time.Millisecond)
		defer cancel()
	}

	wg := sync.WaitGroup{}
	wg.Add(srcCount)
	cl.logger.Debug().Int("src count", srcCount).Msg("wait group value")

	for _, src := range sources {
		cl.logger.Debug().Str("source id", src.ID()).Msg("loading from source")
		go cl.load(ctx, src, serviceNames, resultsCh, &wg)
	}

	cl.logger.Debug().Msg("waiting fo wg")
	wg.Wait()
	cl.logger.Debug().Msg("closing results channel")
	close(resultsCh) // after wg.Wait, so it is ok

	for res := range resultsCh {
		results = append(results, res)
	}

	cl.logger.Debug().Interface("results list", results).Send()
	return results
}
