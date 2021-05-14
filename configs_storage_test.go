package conf_test

import (
	"github.com/golang/mock/gomock"
	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConfigStorage", func() {
	var (
		mockController *gomock.Controller
		configMock     *mocks.MockConfig
		confStorage    conf.ConfigsStorage
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		configMock = mocks.NewMockConfig(mockController)
		confStorage = conf.NewSyncedConfigsStorage()
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("Get,Set", func() {
		serviceName := "testService"

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
