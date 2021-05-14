package sources_test

import (
	"context"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// testConf represents.
type testConf struct{}

func (t testConf) Get(_, _ string) string {
	return ""
}

// Set
func (t testConf) Set(_, _ string) {}

// Fmt.
func (t testConf) Fmt(_ string) (string, error) {
	return "", nil
}

var _ = Describe("Mapsource", func() {
	var (
		priority    = 4
		serviceName = "service1"
		cfgData     = map[string]conf.Config{
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
				const prt = 1
				tmpSrc := sources.NewMapSource(prt, cfgData)
				id := tmpSrc.ID()
				_, ok := ids[id]
				Expect(ok).To(BeFalse())
				ids[id] = struct{}{}
			})
		}
	})

	It("should load without errors", func() {
		err := src.Load(context.Background(), []string{})
		Expect(err).To(BeNil())
	})

	It("should return priority", func() {
		p := src.Priority()
		Expect(p).Should(Equal(priority))
	})

	It("should close without errors", func() {
		err := src.Close(context.Background())
		Expect(err).To(BeNil())
	})

	Describe("getting service config", func() {
		It("should return service config", func() {
			cfg, err := src.ServiceConfig(serviceName)
			Expect(err).To(BeNil())
			Expect(cfg).NotTo(BeNil())
		})

		It("should return service config", func() {
			_, err := src.ServiceConfig("inexisting")
			Expect(err).NotTo(BeNil())
		})
	})
})
