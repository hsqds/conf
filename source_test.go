package conf_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
)

// TestSyncedSourcesStorage
func TestSyncedSourcesStorage(t *testing.T) {
	t.Parallel()

	var (
		srcID = "test-source-id"

		newMockSource = func() (*mocks.MockSource, func()) {
			ctrl := gomock.NewController(t)
			srcMock := mocks.NewMockSource(ctrl)

			cleanup := func() {
				ctrl.Finish()
			}

			return srcMock, cleanup
		}
	)

	t.Run("should append source to storage", func(t *testing.T) {
		t.Parallel()

		storage := conf.NewSyncedSourcesStorage()
		srcMock, cleanup := newMockSource()
		defer cleanup()

		srcMock.EXPECT().ID().Return(srcID).Times(1)

		err := storage.Append(srcMock)
		assert.Nil(t, err)
	})

	t.Run("should return error if source is already in the storage", func(t *testing.T) {
		t.Parallel()

		storage := conf.NewSyncedSourcesStorage()
		srcMock, cleanup := newMockSource()
		defer cleanup()

		srcMock.EXPECT().ID().Return(srcID).Times(2)
		err := storage.Append(srcMock)
		assert.Nil(t, err)

		err = storage.Append(srcMock)
		assert.IsType(t, err, conf.ErrSourceIsNotUnique{})
	})

	t.Run("should return source by id", func(t *testing.T) {
		t.Parallel()

		storage := conf.NewSyncedSourcesStorage()
		srcMock, cleanup := newMockSource()
		defer cleanup()

		srcMock.EXPECT().ID().Return(srcID).Times(1)
		err := storage.Append(srcMock)
		assert.Nil(t, err)

		src, err := storage.ByID(srcID)
		assert.Nil(t, err)
		assert.Equal(t, srcMock, src)
	})

	t.Run("should return error when source not found by id", func(t *testing.T) {
		t.Parallel()

		storage := conf.NewSyncedSourcesStorage()

		_, err := storage.ByID(srcID)
		assert.IsType(t, conf.ErrSourceNotFound{}, err)
	})

	t.Run("should return sources list", func(t *testing.T) {
		t.Parallel()

		var (
			srcID2   = "test-source-id-2"
			storage  = conf.NewSyncedSourcesStorage()
			ctrl     = gomock.NewController(t)
			srcMock  = mocks.NewMockSource(ctrl)
			srcMock2 = mocks.NewMockSource(ctrl)
		)

		defer ctrl.Finish()

		srcMock.EXPECT().ID().Return(srcID).Times(1)
		srcMock2.EXPECT().ID().Return(srcID2).Times(1)

		err := storage.Append(srcMock)
		assert.Nil(t, err)
		err = storage.Append(srcMock2)
		assert.Nil(t, err)

		lst := storage.List()
		assert.Equal(t, []conf.Source{srcMock, srcMock2}, lst)
	})
}
