package sources_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
)

// TestMapSource
func TestMapSource(t *testing.T) {
	t.Parallel()

	const (
		priority = 50
		svc1     = "service1"
		svc2     = "service2"
	)

	t.Run("ID", func(t *testing.T) {
		t.Parallel()

		s := sources.NewMapSource(priority, map[string]conf.Config{})
		assert.Contains(t, s.ID(), "map")
	})

	t.Run("Priority", func(t *testing.T) {
		t.Parallel()

		s := sources.NewMapSource(priority, map[string]conf.Config{})
		assert.Equal(t, s.Priority(), priority)
	})

	t.Run("Load", func(t *testing.T) {
		t.Parallel()

		s := sources.NewMapSource(priority, map[string]conf.Config{})
		err := s.Load(context.Background(), []string{svc1})
		assert.Nil(t, err)
	})

	t.Run("ServiceConfig", func(t *testing.T) {
		t.Parallel()

		var (
			svc1Cfg = conf.NewMapConfig(map[string]string{
				"key1": "val1",
				"key2": "val2",
			})

			d = map[string]conf.Config{
				svc1: svc1Cfg,
			}
		)

		t.Run("should return valid config", func(t *testing.T) {
			t.Parallel()

			s := sources.NewMapSource(priority, d)
			r, err := s.ServiceConfig(svc1)

			assert.Nil(t, err)
			assert.Equal(t, svc1Cfg, r)
		})

		t.Run("should return ServiceConfigError when no config for service", func(t *testing.T) {
			t.Parallel()

			s := sources.NewMapSource(priority, d)
			_, err := s.ServiceConfig(svc2)

			assert.IsType(t, sources.ServiceConfigError{}, err)
		})
	})
}
