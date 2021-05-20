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

	t.Run("should return ServiceConfigStorageError when no config for service", func(t *testing.T) {
		t.Parallel()

		s := conf.NewSyncedConfigsStorage()
		_, err := s.ByServiceName("serviceName")
		assert.IsType(t, conf.ServiceConfigStorageError{}, err)
	})
}
