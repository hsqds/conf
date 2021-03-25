package internal

import "context"

// Provider represents config provider
type Provider struct {
	sources PrioritizedSourcesSet
	loader  *Loader
}

// NewProvider
func NewProvider(sources ...[]PrioritizedSource) *Provider {
	return &Provider{
		sources: PrioritizedSourcesSet(sources),
	}
}

// Load
func (p *Provider) Load(ctx context.Context) (Config, error) {
	data, err := p.loader.Get(ctx, p.sources)
	if err != nil {
		// TODO: loading errors should not be critical
	}

	cfg, err := p.merge(data)
	if err != nil {
		// TODO: loading errors should not be critical

	}
	return cfg, nil
}

// merge
func (p *Provider) merge(cfgLst ConfigsList) (Config, error) {
	return nil, nil
}
