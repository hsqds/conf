package conf_test

import (
	"testing"

	"github.com/hsqds/conf"
	"github.com/stretchr/testify/assert"
)

// TestMapConfig
func TestMapConfig(t *testing.T) {
	t.Parallel()

	const (
		value1 = "value1"
		value2 = "value2"
		defVal = "defaultValue"
	)

	cfgData := map[string]string{
		"key1": value1,
		"key2": value2,
	}

	t.Run("NewMapConfig init with nil data", func(t *testing.T) {
		t.Parallel()

		c := conf.NewMapConfig(nil)
		assert.IsType(t, &conf.MapConfig{}, c)
	})

	t.Run("Get/Set", func(t *testing.T) {
		t.Parallel()

		t.Run("Get should return correct value", func(t *testing.T) {
			t.Parallel()

			const (
				newKey = "newKey"
				newVal = "newVal"
			)

			c := conf.NewMapConfig(cfgData)
			v1 := c.Get("key1", defVal)
			assert.Equal(t, value1, v1)

			c.Set(newKey, newVal)
			nv := c.Get(newKey, defVal)
			assert.Equal(t, newVal, nv)
		})

		t.Run("Get should return default value if key is not set", func(t *testing.T) {
			t.Parallel()

			c := conf.NewMapConfig(cfgData)
			v1 := c.Get("key99", defVal)
			assert.Equal(t, defVal, v1)
		})
	})

	t.Run("Fmt", func(t *testing.T) {
		t.Parallel()

		const (
			pattern = "{{.key1}}-{{.key2}}"
		)

		t.Run("should return correct formatted string", func(t *testing.T) {
			t.Parallel()

			c := conf.NewMapConfig(cfgData)
			exp := value1 + "-" + value2

			r, err := c.Fmt(pattern)
			assert.Nil(t, err)
			assert.Equal(t, exp, r)
		})

		t.Run("should return error if pattern is invalid", func(t *testing.T) {
			t.Parallel()

			c := conf.NewMapConfig(cfgData)
			_, err := c.Fmt(".}}{{")
			assert.Error(t, err)
		})

		t.Run("should return error if pattern rendering failed", func(t *testing.T) {
			t.Parallel()
		})
	})
}
