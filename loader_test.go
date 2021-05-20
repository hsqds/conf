package conf_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
)

// TestConfigsLoader
func TestConfigsLoader(t *testing.T) {
	t.Parallel()

	const (
		configsCount     = 3
		svc1, svc2, svc3 = "service1", "service2", "service3"
	)

	var (
		cfg   = conf.NewMapConfig(map[string]string{})
		srcID = fmt.Sprint(1)

		newSourceMock = func() (*mocks.MockSource, func()) {
			ctrl := gomock.NewController(t)
			src := mocks.NewMockSource(ctrl)

			finish := func() {
				ctrl.Finish()
			}

			return src, finish
		}
	)

	t.Run("should load config", func(t *testing.T) {
		t.Parallel()

		const (
			priority = 50
		)

		services := []string{svc1, svc2, svc3}

		src, finish := newSourceMock()
		defer finish()

		loader := new(conf.ConfigsLoader)

		src.EXPECT().ID().Return(srcID).Times(1)
		src.EXPECT().Load(gomock.Any(), services).Return(nil).Times(1)
		src.EXPECT().ServiceConfig(gomock.Any()).Return(cfg, nil).Times(len(services))
		src.EXPECT().Close(gomock.Any()).Times(1)
		src.EXPECT().Priority().Return(priority).Times(len(services))

		mockSources := []conf.Source{src}

		res := loader.Load(context.Background(), mockSources, services)

		assert.Equal(t, 3, len(res))
		assert.Equal(t, []conf.LoadResult{
			{SourceID: srcID, Config: cfg, Err: nil, Service: svc1, Priority: priority},
			{SourceID: srcID, Config: cfg, Err: nil, Service: svc2, Priority: priority},
			{SourceID: srcID, Config: cfg, Err: nil, Service: svc3, Priority: priority},
		}, res)
	})

	t.Run("LoadResult should contain error when loading from source failed", func(t *testing.T) {
		t.Parallel()

		var (
			err     = errors.New("fail")
			loadErr = conf.LoadError{
				Err:      err,
				Service:  svc1,
				SourceID: srcID,
			}
			services = []string{svc1}
			loader   = new(conf.ConfigsLoader)
		)

		src, finish := newSourceMock()
		defer finish()

		src.EXPECT().ID().Return(srcID).Times(2)
		src.EXPECT().Load(gomock.Any(), services).Return(err).Times(1)
		src.EXPECT().Close(gomock.Any()).Times(1)

		mockSources := []conf.Source{src}

		res := loader.Load(context.Background(), mockSources, services)

		assert.Equal(t, []conf.LoadResult{
			{SourceID: srcID, Config: nil, Err: loadErr},
		}, res)
	})

	t.Run("LoadResult should contain error when source has no config for service", func(t *testing.T) {
		t.Parallel()

		var (
			err      = errors.New("fail")
			loadErr1 = conf.LoadError{
				SourceID: srcID,
				Service:  svc1,
				Err:      err,
			}
			loadErr3 = conf.LoadError{
				SourceID: srcID,
				Service:  svc3,
				Err:      err,
			}
			services = []string{svc1, svc3}
			loader   = new(conf.ConfigsLoader)
		)

		src, finish := newSourceMock()
		defer finish()

		src.EXPECT().ID().Return(srcID).Times(3)
		src.EXPECT().Load(gomock.Any(), services).Return(nil).Times(1)
		src.EXPECT().ServiceConfig(gomock.Any()).Return(nil, err).Times(len(services))
		src.EXPECT().Close(gomock.Any()).Times(1)

		mockSources := []conf.Source{
			conf.Source(src),
		}

		res := loader.Load(context.Background(), mockSources, services)

		assert.Equal(t, []conf.LoadResult{
			{SourceID: srcID, Config: nil, Err: loadErr1, Service: svc1},
			{SourceID: srcID, Config: nil, Err: loadErr3, Service: svc3},
		}, res)
	})
}
