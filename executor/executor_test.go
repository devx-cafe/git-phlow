package executor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/praqma/git-phlow/executor"
)

var _ = Describe("Executor", func() {

	Describe("The RunCommand function", func() {

		Context("called with valid command ls", func() {
			It("should not return an error", func() {
				_, err := RunCommand("git", "--version")

				Ω(err).Should(BeNil())
			})
		})

		Context("called with invalid command", func() {
			It("should fail", func() {
				output, err := RunCommand("lsk", "-lah")
				Ω(output).Should(BeEmpty())
				Ω(err).ShouldNot(BeNil())
			})

		})

	})

	Describe("The RunGit function", func() {

		Context("called with valid command ls", func() {
			It("should not return an error", func() {
				_, err := RunGit("git", "--version")

				Ω(err).Should(BeNil())
			})
		})

		Context("called with invalid command", func() {
			It("should fail", func() {
				output, err := RunGit("lsk", "-lah")
				Ω(output).Should(BeEmpty())
				Ω(err).ShouldNot(BeNil())
			})

		})

	})

})
