package conf_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
	"github.com/hsqds/conf/test/stubs"
)

var _ = Describe("Loader", func() {
	var (
		loader         *conf.ConfigsLoader
		mockController *gomock.Controller
		mockSources    []conf.Source

		src *mocks.MockSource

		svc1, svc2, svc3 = "service1", "service2", "service3"
	)

	BeforeEach(func() {
		loader = conf.NewConfigsLoader(&log.Logger)
		mockController = gomock.NewController(GinkgoT())

		src = mocks.NewMockSource(mockController)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("Load", func() {
		It("should load config", func() {
			cfg := stubs.NewTestConfig()
			configsCount := 3
			srcID := fmt.Sprint(1)
			services := []string{svc1, svc2, svc3}

			src.EXPECT().ID().Return(srcID).Times(2)
			src.EXPECT().Load(gomock.Any(), services).Return(nil).Times(1)
			src.EXPECT().ServiceConfig(gomock.Any()).Return(conf.Config(&cfg), nil).Times(len(services))
			src.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

			mockSources = []conf.Source{
				conf.Source(src),
			}

			res := loader.Load(context.Background(), mockSources, services)
			Expect(len(res)).Should(Equal(configsCount))
			Expect(res).Should(Equal([]conf.LoadResult{
				{SourceID: srcID, Config: &cfg, Err: nil, Service: svc1},
				{SourceID: srcID, Config: &cfg, Err: nil, Service: svc2},
				{SourceID: srcID, Config: &cfg, Err: nil, Service: svc3},
			}))
		})

		It("LoadResult should contain error when loading from source failed", func() {
			srcID := fmt.Sprint(1)
			err := errors.New("fail")
			services := []string{svc1}

			src.EXPECT().ID().Return(srcID).Times(2)
			src.EXPECT().Load(gomock.Any(), services).Return(err).Times(1)
			src.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

			mockSources = []conf.Source{
				conf.Source(src),
			}
			res := loader.Load(context.Background(), mockSources, services)

			Expect(len(res)).Should(Equal(1))
			Expect(res).Should(Equal([]conf.LoadResult{
				{SourceID: srcID, Config: nil, Err: err, Service: ""},
			}))
		})

		It("LoadResult should contain error when service has no config for ", func() {
			srcID := fmt.Sprint(1)
			err := errors.New("fail")
			services := []string{svc1, svc3}
			configsCount := len(services)

			src.EXPECT().ID().Return(srcID).Times(2)
			src.EXPECT().Load(gomock.Any(), services).Return(nil).Times(1)
			src.EXPECT().ServiceConfig(gomock.Any()).Return(nil, err).Times(configsCount)
			src.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

			mockSources = []conf.Source{
				conf.Source(src),
			}

			res := loader.Load(context.Background(), mockSources, services)

			Expect(len(res)).Should(Equal(configsCount))
			Expect(res).Should(Equal([]conf.LoadResult{
				{SourceID: srcID, Config: nil, Err: err, Service: svc1},
				{SourceID: srcID, Config: nil, Err: err, Service: svc3},
			}))
		})
	})
})
