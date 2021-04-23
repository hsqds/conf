package sources_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hate-squids/config-provider/provider"
	"github.com/hate-squids/config-provider/sources"
)

// testConf represents
type testConf struct{}

// funcname
func (t testConf) Get(_ string, _ string) (string, error) {
	return "", nil
}

var _ = Describe("Mapsource", func() {
	var (
		priority    = 4
		serviceName = "service1"
		cfgData     = map[string]provider.Config{
			serviceName: testConf{},
		}
		src = sources.NewMapSource(priority, cfgData)
	)

	Describe("initialization", func() {
		var (
			c   = 10
			ids = make(map[string]struct{}, c)
		)

		for i := 0; i < c; i++ {
			It("should have unique ID", func() {
				tmpSrc := sources.NewMapSource(i, cfgData)
				id := tmpSrc.ID()
				_, ok := ids[id]
				Expect(ok).To(BeFalse())
				ids[id] = struct{}{}
			})
		}
	})

	It("should ping without errors", func() {
		err := src.Ping(context.Background())
		Expect(err).To(BeNil())
	})

	It("should load without errors", func() {
		err := src.Load(context.Background())
		Expect(err).To(BeNil())
	})

	It("should return priority", func() {
		p := src.GetPriority()
		Expect(p).Should(Equal(priority))
	})

	It("should close without errors", func() {
		err := src.Close(context.Background())
		Expect(err).To(BeNil())
	})

	Describe("getting service config", func() {
		It("should return service config", func() {
			cfg, err := src.GetServiceConfig(serviceName)
			Expect(err).To(BeNil())
			Expect(cfg).NotTo(BeNil())
		})

		It("should return service config", func() {
			_, err := src.GetServiceConfig("inexisting")
			Expect(err).NotTo(BeNil())
		})
	})
})