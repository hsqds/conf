package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	defaultLoadTimeout = 1000 // ms
)

// ProviderOption represents provision options
// options may affect subscription and config loading processes
type ProviderOption struct {
	Name  string
	Value string
}

// Provider represents config provider
type ConfigProvider struct {
	sources SourcesStorage
	configs ConfigsStorage
	loader  Loader
}

// NewConfigProvider
func NewConfigProvider(sourcesStorage SourcesStorage, configsStorage ConfigsStorage,
	loader Loader) *ConfigProvider {
	return &ConfigProvider{
		sources: sourcesStorage,
		configs: configsStorage,
		loader:  loader,
	}
}

// GetServiceConfig provide service config from cache
func (p *ConfigProvider) GetServiceConfig(ctx context.Context, serviceName string, opts ...*ProviderOption) (Config, error) {

	return nil, nil
}

// SubscribeForServiceConfig creates a subscription for service
// config updates. Returns "signal channel" of empty structures just
// to notify consumer about updates.
//
// TODO: make sure sending `Config` interface by channel is a bad idea
func (p *ConfigProvider) SubscribeForServiceConfig(ctx context.Context, serviceName string, opts ...*ProviderOption) (chan struct{}, error) {
	return nil, nil
}

// AddSource adds source to source storage
func (p *ConfigProvider) AddSource(src Source) error {
	err := p.sources.Set(src.ID(), src)
	if err != nil {
		return fmt.Errorf("could not set source: %w", err)
	}
	return nil
}

// Load updates services config in cache
func (p *ConfigProvider) Load(ctx context.Context, services ...string) error {
	type configPriority struct {
		cfg Config
		prt int // priority
	}

	ctx, cancel := context.WithTimeout(ctx, defaultLoadTimeout*time.Millisecond)
	defer cancel()

	tmpConfigs := make(map[string]configPriority, len(services))
	var priority int

	for _, result := range p.loader.Load(ctx, p.sources.List(), services) {
		log.Debug().Interface("result", result).Send()
		priority = result.Source.GetPriority()

		cfg, ok := tmpConfigs[result.Service]
		if !ok && priority > cfg.prt {
			tmpConfigs[result.Service] = configPriority{
				cfg: result.Config,
				prt: priority,
			}
		}
	}

	return nil
}
