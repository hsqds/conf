package provider_test

import (
	"context"

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
		prov = provider.NewConfigProvider(ssMock, csMock, loaderMock)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("add source", func() {
		It("should add source", func() {
			srcMock.EXPECT().
				ID().
				Return("id").
				Times(1)

			err := prov.AddSource(srcMock)
			Expect(err).To(BeNil())
		})

		It("should return error when source ID is not unique", func() {
			srcMock.EXPECT().
				ID().
				Return("not_unique").
				Times(2)

			err := prov.AddSource(srcMock)
			Expect(err).To(BeNil())
			err = prov.AddSource(srcMock)
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("get service config", func() {
		var (
			serviceName = "testsvc"
		)

		It("should return service config", func() {
			cfg, err := prov.GetServiceConfig(context.Background(), serviceName)
			Expect(err).To(BeNil())
			Expect(cfg).NotTo(BeNil())
		})
	})

})
