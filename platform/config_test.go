package platform_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/praqma/git-phlow/platform"
	. "github.com/onsi/gomega"
	"github.com/go-errors/errors"
)

var _ = Describe("Config", func() {

	Describe("Testing configurator", func() {
		Context("", func() {
			It("Set should not panic", func() {
				conf := platform.KeyConfiguration{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				var fun = func() { conf.Set(platform.PhlowToken, "value") }
				Ω(fun).ToNot(Panic())
			})

			It("Unset should not panic", func() {
				conf := platform.KeyConfiguration{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				var fun = func() { conf.Set(platform.PhlowToken, "value") }
				Ω(fun).ToNot(Panic())
			})

			It("Get should return EXPECTED", func() {
				EXPECTED := "token"
				conf := platform.KeyConfiguration{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				ACTUAL := conf.Get(platform.PhlowUser)
				Ω(ACTUAL).Should(Equal(EXPECTED))
			})

		})

		Context("when it panics", func() {
			conf := platform.KeyConfiguration{Run: func(command string, argv ...string) (string, error) {
				return "", errors.New("")
			}}

			It("Set should panic", func() {
				var f = func() { conf.Set("invalid", "value") }
				Ω(f).To(Panic())
			})

			It("Get should panic", func() {
				var f = func() { conf.Get("invalid") }
				Ω(f).To(Panic())
			})
			It("Unset should panic", func() {

				var f = func() { conf.Unset("invalid") }
				Ω(f).To(Panic())
			})

		})

	})

})
