package provider_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hate-squids/config-provider/internal/provider"
	"github.com/hate-squids/config-provider/test/mocks"
)

var _ = Describe("ConfigStorage", func() {
	var (
		mockController *gomock.Controller
		configMock     *mocks.MockConfig
		confStorage    provider.ConfigsStorage
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		configMock = mocks.NewMockConfig(mockController)
		confStorage = provider.NewSyncedConfigsStorage()
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("Get,Set", func() {
		var serviceName = "testService"

		It("should get/set configs without errors", func() {
			err := confStorage.Set(serviceName, configMock)
			Expect(err).To(BeNil())
			serviceConf, err := confStorage.Get(serviceName)
			Expect(err).To(BeNil())
			Expect(serviceConf).To(Equal(configMock))
		})

		It("should return error when service config does not exist", func() {
			_, err := confStorage.Get(serviceName)
			Expect(err).NotTo(BeNil())
		})
	})
})
