package conf

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

// LoadResult represents.
type LoadResult struct {
	SourceID string
	Config   Config
	Err      error
	Service  string
	Priority int
}

// LoadError represents
type LoadError struct {
	SourceID string
	Service  string
	Err      error
}

// Error
func (e LoadError) Error() string {
	return fmt.Sprintf(
		"could not load service (%q) config from source (%q): %s",
		e.Service, e.SourceID, e.Err,
	)
}

// Loader.
type Loader interface {
	Load(ctx context.Context, sources []Source, serviceNames []string) []LoadResult
}

// Loader represents.
type ConfigsLoader struct {
	logger zerolog.Logger
}

// NewConfigsLoader.
func NewConfigsLoader(logger *zerolog.Logger) *ConfigsLoader {
	*logger = logger.With().Str("component", "conf.loader").Logger().Level(zerolog.DebugLevel)

	return &ConfigsLoader{*logger}
}

// load call source loading.
func (cl *ConfigsLoader) load(ctx context.Context, src Source,
	services []string, resultsCh chan<- LoadResult, wg *sync.WaitGroup) {
	cl.logger.Debug().Msg("start loading routine")

	defer src.Close(ctx)

	defer wg.Done()

	result := LoadResult{
		SourceID: src.ID(),
	}

	cl.logger.Debug().Msg("Start loading from source")

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
		cl.logger.Debug().Str("service", svc).Msg("getting service config")

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
	cl.logger.Debug().Int("src count", srcCount).Msg("wait group value")

	for _, src := range sources {
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
