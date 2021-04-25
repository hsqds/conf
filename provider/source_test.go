package provider_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hate-squids/config-provider/provider"
	"github.com/hate-squids/config-provider/test/mocks"
)

var _ = Describe("Source", func() {
	var (
		storage        *provider.SyncedSourcesStorage
		mockController *gomock.Controller
		mockSrc        *mocks.MockSource

		key = "srcKey"
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockSrc = mocks.NewMockSource(mockController)
		storage = provider.NewSyncedSourcesStorage()
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("Set", func() {
		It("should set value", func() {
			err := storage.Set(key, mockSrc)
			Expect(err).To(BeNil())
		})

		It("should return error if key is already set", func() {
			err := storage.Set(key, mockSrc)
			Expect(err).To(BeNil())
			err = storage.Set(key, mockSrc)
			Expect(err).NotTo(BeNil())
		})
	})
})
