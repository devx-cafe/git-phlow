package setting_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/code-cafe/git-phlow/setting"
	. "github.com/onsi/gomega"
	"github.com/code-cafe/git-phlow/githandler"
)

var _ = Describe("Setting", func() {

	Describe("LoadProjectSetting", func() {

		Context("default config block with no config", func() {

			It("should not exit but return internal defaults", func() {
				git := githandler.Git{Run: func(git string, sub string, argv ...string) (string, error) {
					return "", nil
				}}
				conf := setting.LoadSettings("phlow", git)
				Ω(conf.DeliveryBranchPrefix).Should(Equal(setting.InternalDefaultDeliveryBranchPrefix))
				Ω(conf.IssueApi).Should(Equal(setting.InternalDefaultApi))
				Ω(conf.Service).Should(Equal(setting.InternalDefaultService))
				Ω(conf.IntegrationBranch).Should(Equal(setting.InternalDefaultIntegrationBranch))
				Ω(conf.Remote).Should(Equal(setting.InternalDefaultRemote))

			})

			Context("default config block with existing fields", func() {
				It("should return the fields and not the internal default", func() {
					git := githandler.Git{Run: func(git string, sub string, argv ...string) (string, error) {
						return "mycustomconfig", nil
					}}
					conf := setting.LoadSettings("phlow", git)
					Ω(conf.DeliveryBranchPrefix).Should(Equal("mycustomconfig"))
				})
			})

		})
	})

	Describe("Test Validation for non-optional params", func() {
		var set = setting.ProjectSetting{}

		BeforeEach(func() {
			set = setting.ProjectSetting{
				Service:              "not empty",
				IntegrationBranch:    "not empty",
				DeliveryBranchPrefix: "not empty",
				Remote:               "not empty",
				IssueApi:             "not empty",
				PipelineUrl:          "not empty"}

		})

		Context("With missing IssueUrl", func() {
			It("Should return error", func() {
				set.IssueApi = ""
				err := setting.ValidateLoadedSetting(&set)
				Ω(err).ShouldNot(BeNil())

			})

		})

		Context("With missing remote", func() {
			It("Should return error", func() {
				set.Remote = ""
				err := setting.ValidateLoadedSetting(&set)
				Ω(err).ShouldNot(BeNil())

			})

		})

		Context("With missing branch prefix", func() {
			It("Should return error", func() {
				set.DeliveryBranchPrefix = ""
				err := setting.ValidateLoadedSetting(&set)
				Ω(err).ShouldNot(BeNil())

			})

		})

		Context("With missing integration branch", func() {
			It("Should return error", func() {
				set.IntegrationBranch = ""
				err := setting.ValidateLoadedSetting(&set)
				Ω(err).ShouldNot(BeNil())
			})
		})

		Context("With missing service", func() {
			It("Should return error", func() {
				set.Service = ""
				err := setting.ValidateLoadedSetting(&set)
				Ω(err).ShouldNot(BeNil())
			})
		})

	})

})
