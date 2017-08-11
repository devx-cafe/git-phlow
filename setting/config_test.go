package setting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"errors"
	"github.com/praqma/git-phlow/setting"
)

var _ = Describe("Config", func() {
	Describe("Testing configurator", func() {
		Context("", func() {
			It("Set should not panic", func() {
				conf := setting.GitConfig{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				var fun = func() { conf.Set(setting.PhlowToken, "value") }
				Ω(fun).ToNot(Panic())
			})

			It("Unset should not panic", func() {
				conf := setting.GitConfig{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				var fun = func() { conf.Set(setting.PhlowToken, "value") }
				Ω(fun).ToNot(Panic())
			})

			It("Get should return EXPECTED", func() {
				EXPECTED := "token"
				conf := setting.GitConfig{Run: func(command string, argv ...string) (string, error) {
					return "token\n", nil
				}}

				ACTUAL := conf.Get(setting.PhlowUser)
				Ω(ACTUAL).Should(Equal(EXPECTED))
			})

		})

		Context("when it panics", func() {
			conf := setting.GitConfig{Run: func(command string, argv ...string) (string, error) {
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
