package provider_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hate-squids/config-provider/internal/provider"
	"github.com/hate-squids/config-provider/test/mocks"
)

var _ = Describe("Source", func() {
	var (
		storage        *provider.SyncedSourcesStorage
		mockController *gomock.Controller
		mockSrc        *mocks.MockSource

		srcID = "test-source-id"
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		storage = provider.NewSyncedSourcesStorage()
		mockSrc = mocks.NewMockSource(mockController)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Describe("Append,Get,List", func() {
		It("should append source to storage", func() {
			mockSrc.EXPECT().ID().Return(srcID).Times(1)

			err := storage.Append(mockSrc)
			Expect(err).To(BeNil())
		})

		It("should return error if source is already in storage", func() {
			mockSrc.EXPECT().ID().Return(srcID).Times(2)

			err := storage.Append(mockSrc)
			Expect(err).To(BeNil())
			err = storage.Append(mockSrc)
			Expect(err).NotTo(BeNil())
		})

		It("should return source by id", func() {
			mockSrc.EXPECT().ID().Return(srcID).Times(1)
			err := storage.Append(mockSrc)
			Expect(err).To(BeNil())

			src, err := storage.Get(srcID)
			Expect(err).To(BeNil())
			Expect(src).To(Equal(mockSrc))
		})

		It("should return error when source not found by id", func() {
			_, err := storage.Get(srcID)
			Expect(err).NotTo(BeNil())
		})

		It("should return sources list", func() {
			var (
				srcID2   = "test-source-id-2"
				mockSrc2 = mocks.NewMockSource(mockController)
			)

			mockSrc.EXPECT().ID().Return(srcID).Times(1)
			mockSrc2.EXPECT().ID().Return(srcID2).Times(1)

			err := storage.Append(mockSrc)
			Expect(err).To(BeNil())
			err = storage.Append(mockSrc2)
			Expect(err).To(BeNil())

			lst := storage.List()
			Expect(len(lst)).To(Equal(2))
		})
	})
})
