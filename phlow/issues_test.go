package phlow_test

import (
	. "github.com/onsi/ginkgo"
	"os"
	. "github.com/onsi/gomega"
	"github.com/praqma/git-phlow/phlow"
	"runtime"
)

var _ = Describe("Issue", func() {

	Describe("GetPager on windows", func() {
		Context("with pager set", func() {
			It("should return testpager", func() {
				EXPECTED := "testpager"

				err := os.Setenv("PAGER", "testpager")
				ACTUAL := phlow.GetPager()

				Ω(ACTUAL).Should(Equal(EXPECTED))
				Ω(err).Should(BeNil())
			})
		})

		Context("with pager unset", func() {
			It("should return more", func() {
				err := os.Unsetenv("PAGER")

				ACTUAL := phlow.GetPager()
				EXPECTED := "more"

				if runtime.GOOS == "windows" {
					Ω(ACTUAL).Should(Equal(EXPECTED))
					Ω(err).Should(BeNil())
				}

			})

		})

	})

	Describe("GetPager on Unix", func() {
		Context("with pager set", func() {
			It("Should return less", func() {
				EXPECTED := "less"

				err := os.Setenv("PAGER", "less")
				ACTUAL := phlow.GetPager()

				Ω(ACTUAL).Should(Equal(EXPECTED))
				Ω(err).Should(BeNil())
			})
		})

		Context("with pager unset", func() {
			It("Should return nothing", func() {
				err := os.Unsetenv("PAGER")

				ACTUAL := phlow.GetPager()
				EXPECTED := ""

				Ω(ACTUAL).Should(Equal(EXPECTED))
				Ω(err).Should(BeNil())

			})
		})

	})

})
