package conf_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
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

	t.Run("LoadError", func(t *testing.T) {
		t.Parallel()

		const (
			sourceID = "sourceid"
			service  = "serviceName"
		)

		err := errors.New("inner error")

		t.Run("Error", func(t *testing.T) {
			t.Parallel()

			e := conf.LoadError{
				SourceID: sourceID,
				Service:  service,
				Err:      err,
			}

			assert.Contains(t, e.Error(), sourceID)
			assert.Contains(t, e.Error(), service)

			assert.Equal(t, err, e.Unwrap())
		})
	})
}
