package conf_test

import (
	"github.com/hsqds/conf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		config *conf.MapConfig

		k1value = "value1"
		k2value = "value2"
	)

	BeforeEach(func() {
		config = &conf.MapConfig{
			"key1": k1value,
			"key2": k2value,
		}
	})

	AfterEach(func() {
		config = nil
	})

	Describe("Get", func() {
		It("should return value when it is set", func() {
			v1 := config.Get("key1", "default")
			Expect(v1).To(Equal(k1value))

			v2 := config.Get("key2", "default")
			Expect(v2).To(Equal(k2value))
		})

		It("should return default value when no value is set", func() {
			const def = "default"
			v := config.Get("key", def)
			Expect(v).To(Equal(def))
		})
	})

	Describe("Fmt", func() {
		It("should return formatted string", func() {
			f := k1value + "-" + k2value
			s, err := config.Fmt("{{.key1}}-{{.key2}}")
			Expect(err).To(BeNil())
			Expect(s).To(Equal(f))
		})

		It("should return error when pattern is not correct", func() {
			_, err := config.Fmt(".}}{{")
			Expect(err).NotTo(BeNil())
		})
	})
})
