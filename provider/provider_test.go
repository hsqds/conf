package provider_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hate-squids/config-provider/provider"
	"github.com/hate-squids/config-provider/test/mocks"
)

var _ = Describe("Provider", func() {
	var (
		mockController *gomock.Controller
		srcMock        *mocks.MockSource
		ssMock         *mocks.MockSourcesStorage
		csMock         *mocks.MockConfigStorage
		loaderMock     *mocks.MockLoader
		prov           *provider.ConfigProvider
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		srcMock = mocks.NewMockSource(mockController)
		ssMock = mocks.NewMockSourcesStorage(mockController)
		csMock = mocks.NewMockConfigStorage(mockController)
		loaderMock = mocks.NewMockLoader(mockController)
		prov = provider.NewConfigProvider(ssMock, csMock, loaderMock)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("AddSource", func() {
		var (
			srcID = "id"
		)

		It("should add source", func() {
			srcMock.EXPECT().
				ID().
				Return(srcID).
				Times(1)

			ssMock.EXPECT().
				Set(srcID, srcMock).
				Return(nil).
				Times(1)

			err := prov.AddSource(srcMock)
			Expect(err).To(BeNil())
		})

		It("should return error when source set failed", func() {
			srcMock.EXPECT().
				ID().
				Return(srcID).
				Times(1)

			ssMock.EXPECT().
				Set(srcID, srcMock).
				Return(errors.New("fail"))

			err := prov.AddSource(srcMock)
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("get service config", func() {
		var (
			serviceName = "testsvc"
		)

		It("should return service config", func() {
			_ = serviceName
		})
	})

})
