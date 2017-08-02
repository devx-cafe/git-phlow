package platform_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/praqma/git-phlow/platform"
	. "github.com/onsi/gomega"
	"github.com/praqma/git-phlow/executor"
	"github.com/go-errors/errors"
)

var _ = Describe("Platform", func() {

	Describe("Running platform", func() {
		It("should return github", func() {
			pl := platform.Platform{Run: func(command string, argv ...string) (string, error) {
				return "git@github.com:Praqma/git-phlow.git", nil
			}}

			Expected := platform.GITHUB
			Actual := pl.Service()

			Ω(Actual).Should(Equal(Expected))
		})

		It("should return bitbucket", func() {
			pl := platform.Platform{Run: func(command string, argv ...string) (string, error) {
				return "git@bitbucket:Praqma/git-phlow.git", nil
			}}

			Expected := platform.BITBUCKET
			Actual := pl.Service()

			Ω(Actual).Should(Equal(Expected))
		})

		It("should return unrecognized platfrom", func() {
			pl := platform.Platform{Run: func(command string, argv ...string) (string, error) {
				return "jklsjdklasjdlsakjdklas", nil
			}}

			Expected := platform.UNRECOGNIZEDPLATFORM
			Actual := pl.Service()

			Ω(Actual).Should(Equal(Expected))
		})

		It("should return unrecognized platfrom", func() {
			platform := platform.Platform{Run: func(command string, argv ...string) (string, error) {
				return "", errors.New("crash and burn")
			}}

			Ω(func() { platform.Service() }).To(Panic())
		})

		It("default should provide executor", func() {
			platform := platform.NewDefaultPlatform()

			Ω(platform.Run).Should(BeAssignableToTypeOf(executor.Run))

		})

	})

	Describe("", func() {

	})

})
