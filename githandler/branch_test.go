package githandler

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/praqma/git-phlow/testfixture"
)

var _ = Describe("Branch", func() {

	BeforeEach(func() {

		//Runs before each "It" block
		testfixture.CreateTestRepositoryNoLog(false)
	})

	Describe("Branch Function execution", func() {

		It("should return a list of branches", func() {
			info, err := Branch()
			Ω(len(info.List)).Should(Equal(11))
			Ω(err).Should(BeNil())
		})

		It("branch should return Current branch", func() {
			info, err := Branch()
			Ω(info.Current).Should(Equal("master"))
			Ω(err).Should(BeNil())
		})

	})

	Describe("Delete Branch Function", func() {

		It("should delete local branch and return message", func() {
			output, err := BranchDelete("delivered/1-issue-branch", "", false, false)
			info, _ := Branch()

			Ω(err).Should(BeNil())
			Ω(output).ShouldNot(BeEmpty())
			Ω(info.List).Should(HaveLen(10))
		})

		It("should delete remote branch and return message", func() {
			_, err1 := BranchDelete("delivered/24-issue-branch", "origin", true, false)
			_, err2 := BranchDelete("delivered/42-issue-branch", "origin", true, false)
			info, _ := Branch()

			Ω(err1).Should(BeNil())
			Ω(err2).Should(BeNil())

			Ω(info.List).Should(HaveLen(9))
		})

	})

	Describe("Executing BranchDelivered", func() {
		It("should return lists of delivered branches", func() {
			locals, remotes := BranchDelivered("origin")
			Ω(locals).Should(HaveLen(1))
			Ω(remotes).Should(HaveLen(2))
			Ω(remotes).Should(ContainElement("delivered/24-issue-branch"))
		})

	})

	Describe("Running BranchDelivered", func() {

		It("should return list of ready branches", func() {
			remotes := BranchReady("origin", "ready/")

			Ω(remotes).Should(HaveLen(2))
			Ω(remotes).Should(ContainElement("origin/ready/99-issue-branch"))
		})

	})

	Describe("Running tests on 'BranchTime' function", func() {

		It("Should get unix timestamp", func() {
			output, err := BranchTime("origin/ready/99-issue-branch")
			Ω(err).Should(BeNil())
			Ω(output).Should(BeNumerically(">", 100000))

		})

		It("Should fail geting unix timestamp", func() {
			output, err := BranchTime("bluarh.. not a branch")
			Ω(err).ShouldNot(BeNil())
			Ω(output).Should(Equal(-1))
		})

	})

	Describe("Executing BranchRemote", func() {

		It("should return origin/master", func() {
			output, err := branchRemote()

			Ω(err).Should(BeNil())
			Ω(output).Should(Equal("origin/master"))

		})

	})

	AfterEach(func() {
		testfixture.RemoveTestRepositoryNoLog()
	})

})
