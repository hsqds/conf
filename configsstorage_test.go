package conf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
)

// TestConfigsStorage
func TestConfigsStorage(t *testing.T) {
	t.Parallel()

	t.Run("should get and set config", func(t *testing.T) {
		t.Parallel()

		const svc1 = "service1"

		cfg := conf.NewMapConfig(map[string]string{})
		s := conf.NewSyncedConfigsStorage()

		err := s.Set(svc1, cfg)
		assert.Nil(t, err)

		r, err := s.ByServiceName(svc1)
		assert.Nil(t, err)

		assert.Equal(t, cfg, r)
	})

	t.Run("should return ErrServiceConfigNotFound when no config for service", func(t *testing.T) {
		t.Parallel()

		s := conf.NewSyncedConfigsStorage()
		_, err := s.ByServiceName("serviceName")
		assert.IsType(t, conf.ErrServiceConfigNotFound{}, err)
	})

	t.Run("should return true if config exists at the storage and false otherwise", func(t *testing.T) {
		t.Parallel()

		const (
			existingSvc   = "service1"
			inexistingSvc = "inexisting"
		)

		cfg := conf.NewMapConfig(map[string]string{})
		s := conf.NewSyncedConfigsStorage()

		err := s.Set(existingSvc, cfg)
		assert.Nil(t, err)

		ok := s.Has(existingSvc)
		assert.True(t, ok)

		ok = s.Has(inexistingSvc)
		assert.False(t, ok)
	})
}
