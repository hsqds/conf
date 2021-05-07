package conf

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

// Option represents provision options
// options may affect subscription and config loading processes
type Option struct {
	Name  string
	Value string
}

// Provider represents config provider
type ConfigProvider struct {
	sources SourcesStorage
	configs ConfigsStorage
	loader  Loader
	logger  zerolog.Logger
}

// NewDefaultConfigProvider
func NewDefaultConfigProvider(logger *zerolog.Logger) *ConfigProvider {
	return NewConfigProvider(
		NewSyncedSourcesStorage(),
		NewSyncedConfigsStorage(),
		NewConfigsLoader(logger),
		logger,
	)
}

// NewConfigProvider
func NewConfigProvider(sourcesStorage SourcesStorage, configsStorage ConfigsStorage,
	loader Loader, logger *zerolog.Logger) *ConfigProvider {
	return &ConfigProvider{
		sources: sourcesStorage,
		configs: configsStorage,
		loader:  loader,
		logger:  *logger,
	}
}

// ServiceConfig provide service config from cache
func (p *ConfigProvider) ServiceConfig(serviceName string, opts ...*Option) (Config, error) {
	return p.configs.Get(serviceName)
}

// SubscribeForServiceConfig creates a subscription for service
// config updates. Returns channel of Configs
func (p *ConfigProvider) SubscribeForServiceConfig(ctx context.Context, serviceName string, opts ...*Option) (chan Config, error) {
	return nil, nil
}

// AddSource adds source to source storage
func (p *ConfigProvider) AddSource(src Source) error {
	err := p.sources.Append(src)
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

	tmpConfigs := make(map[string]configPriority, len(services))
	var priority int

	for _, result := range p.loader.Load(ctx, p.sources.List(), services) {
		p.logger.Debug().Interface("result", result)

		if result.Err != nil {
			p.logger.Warn().Err(result.Err).Send()
			continue
		}

		src, err := p.sources.Get(result.SourceID)
		if err != nil {
			err = fmt.Errorf("could not get source by id (%q): %w", result.SourceID, err)
			p.logger.Error().Err(err).Send()
		}

		priority = src.Priority()

		cfgP, ok := tmpConfigs[result.Service]
		if !ok || priority > cfgP.prt {
			tmpConfigs[result.Service] = configPriority{
				cfg: result.Config,
				prt: priority,
			}
		}
	}

	for svcName, cfgP := range tmpConfigs {
		p.logger.Debug().Str("service name", svcName).Interface("config", cfgP.cfg).Msg("updating config cache")
		err := p.configs.Set(svcName, cfgP.cfg)
		if err != nil {
			return fmt.Errorf("could not update configs storage: %w", err)
		}
	}

	return nil
}

// Close
func (p *ConfigProvider) Close(ctx context.Context) {
	for _, src := range p.sources.List() {
		src.Close(ctx)
	}
}
