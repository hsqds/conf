package conf_test

import (
	"context"
	"errors"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
)

var _ = Describe("Provider", func() {
	var (
		mockController *gomock.Controller
		srcMock        *mocks.MockSource
		ssMock         *mocks.MockSourcesStorage
		csMock         *mocks.MockConfigStorage
		loaderMock     *mocks.MockLoader
		prov           *conf.ConfigProvider

		logger = zerolog.New(os.Stderr)
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		srcMock = mocks.NewMockSource(mockController)
		ssMock = mocks.NewMockSourcesStorage(mockController)
		csMock = mocks.NewMockConfigStorage(mockController)
		loaderMock = mocks.NewMockLoader(mockController)
		prov = conf.NewConfigProvider(ssMock, csMock, loaderMock, &logger)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("AddSource", func() {
		It("should add source", func() {
			ssMock.EXPECT().Append(srcMock).Return(nil).Times(1)

			err := prov.AddSource(srcMock)
			Expect(err).To(BeNil())
		})

		It("should return error when source set failed", func() {
			ssMock.EXPECT().Append(srcMock).Return(errors.New("fail")).Times(1)

			err := prov.AddSource(srcMock)
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("Load,GetServiceConfig", func() {
		var (
			services = []string{"testsvc1", "testsvc2", "testsvc3"}
			sources  = []conf.Source{
				srcMock,
			}
		)

		It("should load service configs from sources", func() {
			ssMock.EXPECT().List().Return(sources).Times(1)
			loaderMock.EXPECT().Load(gomock.Any(), gomock.Any(), services).
				Return(nil).Times(1)

			err := prov.Load(context.Background(), services...)
			Expect(err).To(BeNil())
		})
	})
})
