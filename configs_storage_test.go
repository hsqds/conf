package conf_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hsqds/conf"
	"github.com/hsqds/conf/test/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
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

// TestConfigsStorage
func TestConfigsStorage(t *testing.T) {
	t.Parallel()

	t.Run("should get and set config", func(t *testing.T) {
		t.Parallel()

		const svc1 = "service1"

		cfg := conf.NewMapConfig(map[string]string{})
		s := conf.NewSyncedConfigsStorage()

		err := s.Set(svc1, cfg)
		assert.Nil(t, err)

		r, err := s.Get(svc1)
		assert.Nil(t, err)

		assert.Equal(t, cfg, r)
	})

	t.Run("should return ServiceConfigStorageError when no config for service", func(t *testing.T) {
		t.Parallel()

		s := conf.NewSyncedConfigsStorage()
		_, err := s.Get("serviceName")
		assert.IsType(t, conf.ServiceConfigStorageError{}, err)
	})
}
