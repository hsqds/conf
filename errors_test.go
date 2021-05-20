package conf_test

import (
	"testing"

	"github.com/hsqds/conf"
	"github.com/stretchr/testify/assert"
)

// TestStoragesErrors
func TestStoragesErrors(t *testing.T) {
	t.Parallel()
	t.Run("ServiceConfigStorageError", func(t *testing.T) {
		t.Parallel()

		const svc1 = "serviceName1"

		e := conf.ServiceConfigStorageError{svc1}

		assert.Contains(t, e.Error(), svc1)
	})

	t.Run("SourceUniquenessError", func(t *testing.T) {
		t.Parallel()

		const sourceID = "sourceID1"

		e := conf.SourceUniquenessError{sourceID}

		assert.Contains(t, e.Error(), sourceID)
	})

	t.Run("SourceStorageError", func(t *testing.T) {
		t.Parallel()

		const sourceID = "sourceID1"

		e := conf.SourceStorageError{sourceID}

		assert.Contains(t, e.Error(), sourceID)
	})
}
