package setting_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/praqma/git-phlow/setting"
	. "github.com/onsi/gomega"
	"github.com/go-errors/errors"
)

var _ = Describe("Setting", func() {

	Describe("Test Local and Global", func() {

		It("local should not be empty", func() {
			global := setting.GetLocal()
			Ω(global).ShouldNot(BeEmpty())
		})

		It("global should not be empty", func() {
			local := setting.GetGlobal()
			Ω(local).ShouldNot(BeEmpty())
		})

	})

	Describe("LoadToolSetting", func() {

		It("should return user and token", func() {
			i := 0
			set := setting.LoadToolSettings(func(command string, argv ...string) (string, error) {
				i++
				if i == 2 {
					return "token", nil
				}
				return "user", nil
			})
			Ω(set.Token).Should(Equal("token"))
			Ω(set.User).Should(Equal("user"))
		})

		It("should panic", func() {
			panicFunc := func() {
				setting.LoadToolSettings(func(command string, argv ...string) (string, error) {
					return "", errors.New("I will trigger a panic")
				})
			}
			Ω(panicFunc).To(Panic())
		})

	})

	Describe("LoadProjectSetting", func() {

		It("should return find local default", func() {
			conf := setting.LoadProjectSettings(setting.GetLocal(), "", "default")
			Ω(conf.File).Should(Equal(".phlow"))
			Ω(conf.IssueURL).Should(Equal("https://api.github.com"))
		})

		It("should find local jira", func() {
			conf := setting.LoadProjectSettings(setting.GetLocal(), "", "jira")
			Ω(conf.File).Should(Equal(".phlow"))
			Ω(conf.IssueURL).Should(Equal("jira"))
		})

		It("no params should set default", func() {
			conf := setting.LoadProjectSettings(setting.GetLocal(), "", "")
			Ω(conf.File).Should(Equal(".phlow"))
			Ω(conf.IssueURL).Should(Equal("https://api.github.com"))
		})

		It("no config files should use internal default", func() {
			conf := setting.LoadProjectSettings("", "", "")
			Ω(conf.File).Should(Equal("none"))
			Ω(conf.Scope).Should(Equal("internal"))
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
				IssueURL:             "not empty",
				PipelineUrl:          "not empty"}

		})

		Context("With missing IssueUrl", func() {
			It("Should return error", func() {
				set.IssueURL = ""
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
