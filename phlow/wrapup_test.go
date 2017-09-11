package phlow_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/praqma/git-phlow/phlow"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wrapup", func() {

	Describe("Testing getting name and issue", func() {

		Context("for github branches", func() {

			It("for github should return error", func() {
				_, err := phlow.GetJIRAIssue("1-hello-issue")
				Ω(err).ShouldNot(BeNil())
			})

			It("for master should return error", func() {
				_, err := phlow.GetJIRAIssue("master")
				Ω(err).ShouldNot(BeNil())
			})

		})

		Context("for jira branches", func() {

			It("should return TIS-46", func() {
				issue, err := phlow.GetJIRAIssue("TIS-46-update-localtransportcontroller-to-handle-multiple-travel-providers-in-one-reservation")
				Ω(err).Should(BeNil())
				Ω(issue).Should(Equal("TIS-46"))
			})

			It("should return TIS-45", func() {
				issue, err := phlow.GetJIRAIssue("TIS-45-email-non-registered-users-to-sign-up-with-teams-in-space")
				Ω(err).Should(BeNil())
				Ω(issue).Should(Equal("TIS-45"))
			})

			It("should return ADO-70", func() {
				issue, err := phlow.GetJIRAIssue("ADO-70-booking-button-randomly-dissapears")
				Ω(err).Should(BeNil())
				Ω(issue).Should(Equal("ADO-70"))
			})

		})

	})
})
