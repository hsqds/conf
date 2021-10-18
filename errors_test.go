package conf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
)

// TestStoragesErrors
func TestStoragesErrors(t *testing.T) {
	t.Parallel()
	t.Run("ErrServiceConfigNotFound", func(t *testing.T) {
		t.Parallel()

		const svc1 = "serviceName1"

		e := conf.ErrServiceConfigNotFound{svc1}

		assert.Contains(t, e.Error(), svc1)
	})

	t.Run("ErrSourceIsNotUnique", func(t *testing.T) {
		t.Parallel()

		const sourceID = "sourceID1"

		e := conf.ErrSourceIsNotUnique{sourceID}

		assert.Contains(t, e.Error(), sourceID)
	})

	t.Run("ErrSourceNotFound", func(t *testing.T) {
		t.Parallel()

		const sourceID = "sourceID1"

		e := conf.ErrSourceNotFound{sourceID}

		assert.Contains(t, e.Error(), sourceID)
	})
}
