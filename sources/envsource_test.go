package sources_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
	"github.com/stretchr/testify/assert"
)

// TestEnvSource
func TestEnvSource(t *testing.T) {
	t.Parallel()

	const (
		prt  = 50
		svc1 = "service1"
		svc2 = "service2"
	)

	t.Run("ID", func(t *testing.T) {
		t.Parallel()

		src := sources.NewEnvSource(prt, []string{})
		id := src.ID()
		if !strings.Contains(id, "env") {
			t.Errorf("source ID should contain 'env' string")
		}
	})

	t.Run("Priority", func(t *testing.T) {
		t.Parallel()
		src := sources.NewEnvSource(prt, []string{})
		if prt != src.Priority() {
			t.Errorf("wrong priority value")
		}
	})

	t.Run("Load/ServiceConfig", func(t *testing.T) {
		var (
			input = []string{
				"service1_key_num1=val11",
				"service1_KEY2=VAL12",
				"SERVICE1_key3=val13",
				"SerVice1_Key4=val14",
				"service1-key5=val15",
				"service1_key6",
				"service2_key1=val21",
				"service2_key2=val22",
				"service3_key1=val31",
				"service3_key2=val32",
			}
			exp = conf.MapConfig{
				"keyNum1": "val11",
				"key2":    "VAL12",
				"key3":    "val13",
				"key4":    "val14",
			}
			svc2Exp = conf.MapConfig{
				"key1": "val21",
				"key2": "val22",
			}
		)

		t.Run("ServiceConfig should return correct config", func(t *testing.T) {
			t.Parallel()

			src := sources.NewEnvSource(prt, input)
			src.Load(context.Background(), []string{svc1, svc2})

			cfg, err := src.ServiceConfig(svc1)
			if err != nil {
				t.Errorf("ServiceConfig should not return error")
			}

			if !assert.Equal(t, exp, cfg) {
				t.Errorf("expect to get %#v got %#v", exp, cfg)
			}
		})

		t.Run("ServiceConfig result should be idempotant", func(t *testing.T) {
			t.Parallel()

			src := sources.NewEnvSource(prt, input)
			src.Load(context.Background(), []string{svc1, svc2, svc1})

			cfg, err := src.ServiceConfig(svc1)
			assert.Nil(t, err)
			assert.Equal(t, exp, cfg)

			cfg, err = src.ServiceConfig(svc2)
			assert.Nil(t, err)
			assert.Equal(t, svc2Exp, cfg)

			cfg, err = src.ServiceConfig(svc1)
			assert.Nil(t, err)
			assert.Equal(t, exp, cfg)
		})

		t.Run("ServiceConfig should return ServiceConfigError when service config not found", func(t *testing.T) {
			t.Parallel()

			src := sources.NewEnvSource(prt, input)
			src.Load(context.Background(), []string{svc1})

			notLoadedService := "serviceName"

			expErr := sources.ServiceConfigError{notLoadedService, src.ID()}

			_, err := src.ServiceConfig(notLoadedService)
			assert.Equal(t, expErr, err)
		})
	})
}
