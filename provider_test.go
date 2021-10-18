package conf_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
)

// TestConfigProvider
func TestConfigProvider(t *testing.T) {
	t.Parallel()

	var ()

	t.Run("AddSource", func(t *testing.T) {
		t.Parallel()

		t.Run("should add source", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				srcMock    = mocks.NewMockSource(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
			)

			defer ctrl.Finish()

			ssMock.EXPECT().Append(srcMock).Return(nil).Times(1)

			err := prov.AddSource(srcMock)
			assert.Nil(t, err)
		})

		t.Run("should return error when source set failed", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				srcMock    = mocks.NewMockSource(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
			)

			ssMock.EXPECT().Append(srcMock).Return(errors.New("fail")).Times(1)

			err := prov.AddSource(srcMock)
			assert.Error(t, err)
		})
	})

	t.Run("Load,ServiceConfig", func(t *testing.T) {
		t.Parallel()

		t.Run("should load service configs from sources", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				srcMock    = mocks.NewMockSource(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
				services   = []string{"testsvc1", "testsvc2", "testsvc3"}
				sources    = []conf.Source{
					srcMock,
				}
			)

			ssMock.EXPECT().List().Return(sources).Times(1)
			loaderMock.EXPECT().Load(gomock.Any(), gomock.Any(), services).Return(nil).Times(1)

			err := prov.Load(context.Background(), services...)
			assert.Nil(t, err)
		})

		t.Run("ServiceConfig should return service config", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
				svc1       = "service1"
				cfg1       = mocks.NewMockConfig(ctrl)
			)

			csMock.EXPECT().ByServiceName(svc1).Return(cfg1, nil).Times(1)
			csMock.EXPECT().Has(svc1).Return(true).Times(1)
			rc1, err := prov.ServiceConfig(svc1)

			assert.Nil(t, err)
			assert.Equal(t, cfg1, rc1)
		})

		t.Run("ServiceConfig should return error when configs storage return error", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
				svc1       = "service1"
			)

			csMock.EXPECT().ByServiceName(svc1).Return(nil, errors.New("fail")).Times(1)
			csMock.EXPECT().Has(svc1).Return(false).Times(1)
			_, err := prov.ServiceConfig(svc1)

			assert.Error(t, err)
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Parallel()

		t.Run("should call every source's Close method", func(t *testing.T) {
			t.Parallel()

			var (
				ctrl       = gomock.NewController(t)
				ssMock     = mocks.NewMockSourcesStorage(ctrl)
				csMock     = mocks.NewMockConfigStorage(ctrl)
				loaderMock = mocks.NewMockLoader(ctrl)
				src1Mock   = mocks.NewMockSource(ctrl)
				src2Mock   = mocks.NewMockSource(ctrl)
				prov       = conf.NewConfigProvider(ssMock, csMock, loaderMock)
			)

			src1Mock.EXPECT().Close(gomock.Any()).Times(1)
			src2Mock.EXPECT().Close(gomock.Any()).Times(1)
			ssMock.EXPECT().List().Return([]conf.Source{src1Mock, src2Mock})

			prov.Close(context.Background())
		})
	})
}
