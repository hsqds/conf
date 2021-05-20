package sources_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
)

// TestFlagsSource
func TestFlagsSource(t *testing.T) {
	t.Parallel()

	const (
		priority = 50
		prefix   = "--"
		svc1     = "service1"
		svc2     = "service2"
	)

	t.Run("ID", func(t *testing.T) {
		t.Parallel()

		src := sources.InitFlagsSource(priority, prefix, []string{})
		assert.Contains(t, src.ID(), "flags")
	})

	t.Run("Priority", func(t *testing.T) {
		t.Parallel()

		src := sources.InitFlagsSource(priority, prefix, []string{})
		assert.Equal(t, priority, src.Priority())
	})

	t.Run("Load/ServiceConfig", func(t *testing.T) {
		t.Parallel()

		var (
			input = []string{
				"--service1-key-num1=val11",
				"--service1-KEY2=VAL12",
				"--sERVICE1-key3=val13",
				"--serVice1-Key4=val14",
				"-service1-key5=val15",
				"--service1-key6",
				"--service2-key1=val21",
				"--service2-key2=val22",
				"--service3-key1=val31",
				"--service3-key2=val32",
			}
			exp = conf.NewMapConfig(map[string]string{
				"keyNum1": "val11",
				"key2":    "VAL12",
				"key3":    "val13",
				"key4":    "val14",
			})
		)

		t.Run("ServiceConfig should return valid config", func(t *testing.T) {
			t.Parallel()

			e := sources.InitFlagsSource(priority, prefix, input)
			e.Load(context.Background(), []string{svc1, svc2})

			cfg, err := e.ServiceConfig(svc1)
			assert.Nil(t, err)
			assert.Equal(t, exp, cfg)
		})

		t.Run("ServiceConfig should return ServiceConfigError when no config loaded for service", func(t *testing.T) {
			t.Parallel()

			e := sources.InitFlagsSource(priority, prefix, input)
			e.Load(context.Background(), []string{svc1})

			_, err := e.ServiceConfig(svc2)
			assert.IsType(t, sources.ServiceConfigError{}, err)
		})
	})
}
