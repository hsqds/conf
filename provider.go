package conf

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// Option represents provision options
// options may affect subscription and config loading processes.
type Option struct {
	Name  string
	Value string
}

// Provider represents config provider.
type ConfigProvider struct {
	sources            SourcesStorage
	configs            ConfigsStorage
	loader             Loader
	defaultLoadTimeout time.Duration
}

// NewDefaultConfigProvider.
func NewDefaultConfigProvider() *ConfigProvider {
	return NewConfigProvider(
		NewSyncedSourcesStorage(),
		NewSyncedConfigsStorage(),
		new(ConfigsLoader),
	)
}

// NewConfigProvider.
func NewConfigProvider(sourcesStorage SourcesStorage, configsStorage ConfigsStorage,
	loader Loader) *ConfigProvider {
	return &ConfigProvider{
		sources:            sourcesStorage,
		configs:            configsStorage,
		loader:             loader,
		defaultLoadTimeout: time.Second,
	}
}

// ServiceConfig provide service config from inner cache.
func (p *ConfigProvider) ServiceConfig(serviceName string, opts ...*Option) (Config, error) {
	var (
		autoload    bool
		loadTimeout = p.defaultLoadTimeout
	)

	for _, opt := range opts {
		switch opt.Name {
		case "autoload":
			v, err := strconv.ParseBool(opt.Value)
			if err != nil {
				return nil, fmt.Errorf("could not parse option value: %w", err)
			}

			autoload = v
		case "loadTimeout":
			v, err := strconv.Atoi(opt.Value)
			if err != nil {
				return nil, fmt.Errorf("could not parse loadTimeout option value: %w", err)
			}

			loadTimeout = time.Duration(v)
		default:
			// TODO: process unknown option name
		}
	}

	if !p.configs.Has(serviceName) && autoload {
		ctx, cancel := context.WithTimeout(context.TODO(), loadTimeout)
		defer cancel()

		p.Load(ctx, serviceName)
	}

	cfg, err := p.configs.ByServiceName(serviceName)
	if err != nil {
		return nil, fmt.Errorf("could not get service config: %w", err)
	}

	return cfg, nil
}

// SubscribeForServiceConfig creates a subscription for service
// config updates. Returns channel of Configs.
func (p *ConfigProvider) SubscribeForServiceConfig(ctx context.Context, serviceName string,
	opts ...*Option) (chan Config, error) {
	return nil, nil
}

// AddSource adds source to source storage.
func (p *ConfigProvider) AddSource(src Source) error {
	if err := p.sources.Append(src); err != nil {
		return fmt.Errorf("could not set source: %w", err)
	}

	return nil
}

// Load updates inner services config cache.
func (p *ConfigProvider) Load(ctx context.Context, services ...string) (loadErrors []LoadError) {
	type configPriority struct {
		cfg Config
		prt int // priority
	}

	tmpConfigs := make(map[string]configPriority, len(services))

	var priority int

	results := p.loader.Load(ctx, p.sources.List(), services)
	for i := range results {
		result := results[i]

		if result.Err != nil {
			loadErrors = append(loadErrors, result.Err.(LoadError))

			continue
		}

		priority = result.Priority

		// TODO: move this logic to []LoadResult
		// Getting set configs. Each config is most prioritized
		// for it's service
		cfgP, ok := tmpConfigs[result.Service]
		if !ok || priority > cfgP.prt {
			tmpConfigs[result.Service] = configPriority{
				cfg: result.Config,
				prt: priority,
			}
		}
	}

	for svcName, cfgP := range tmpConfigs {
		err := p.configs.Set(svcName, cfgP.cfg)
		if err != nil {
			loadErrors = append(loadErrors, LoadError{
				Service:  svcName,
				SourceID: "",
				Err:      fmt.Errorf("could not update configs storage: %w", err),
			})

			continue
		}
	}

	return loadErrors
}

// Close.
func (p *ConfigProvider) Close(ctx context.Context) {
	for _, src := range p.sources.List() {
		src.Close(ctx)
	}
}
