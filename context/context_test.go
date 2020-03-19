package context_test

import (
	"testing"

	. "github.com/code-cafe/git-phlow/context"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var gitBranchOutput = `
  315-github-api-deprecation-notice-for-authentication-via-url-query-parameters
* 327-update-workon-command
  delivered/326-update-readme-with-new-introduction
  master
  hello-world
`

func TestExecutor(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Executor Suite")
}

var _ = ginkgo.Describe("context", func() {

	ginkgo.Context("Test GetBranches", func() {
		ginkgo.It("should return all branches", func() {
			branches := GetBranches(gitBranchOutput)

			Ω(len(branches)).Should(Equal(5))
		})
	})

	ginkgo.Context("Test GetDelivered", func() {
		ginkgo.It("should return all delivered", func() {
			branches := GetBranches(gitBranchOutput)
			delivered := GetDelivered(branches)

			Ω(len(delivered)).Should(Equal(1))
		})
	})

	ginkgo.Context("Test GetGetWorkspaces", func() {
		ginkgo.It("should return all delivered", func() {
			branches := GetBranches(gitBranchOutput)
			workspaces := GetWorkSpaces(branches)

			Ω(len(workspaces)).Should(Equal(2))
		})
	})

	ginkgo.Context("Get organization and repo github uri", func() {
		ginkgo.It("should return organisation and repository with https url", func() {
			org, repo := GetOrganizationAndRepository("https://github.com/code-cafe/phlow-test-strategy.git")

			Ω(org).Should(Equal("code-cafe"))
			Ω(repo).Should(Equal("phlow-test-strategy"))
		})

		ginkgo.It("should return organisation and repository with ssh", func() {
			org, repo := GetOrganizationAndRepository("git@github.com:Org/some-repo.git")

			Ω(org).Should(Equal("Org"))
			Ω(repo).Should(Equal("some-repo"))
		})

		ginkgo.It("ssh remote url with dot should return", func() {
			org, repo := GetOrganizationAndRepository("git@github.com:Praqma/praqma.com.git")

			Ω(org).Should(Equal("Praqma"))
			Ω(repo).Should(Equal("praqma.com"))
		})
	})

	ginkgo.Context("Get organization and repo bitbucket uri", func() {

		ginkgo.It("should return organisation and repository with https url", func() {
			org, repo := GetOrganizationAndRepository("https://bitbucket.com/Org/sOme--repo.git")

			Ω(org).Should(Equal("Org"))
			Ω(repo).Should(Equal("sOme--repo"))
		})

		ginkgo.It("should return organisation and repository with ssh", func() {
			org, repo := GetOrganizationAndRepository("git@bitbucket.com:Org/some-repo.git")

			Ω(org).Should(Equal("Org"))
			Ω(repo).Should(Equal("some-repo"))
		})

		ginkgo.It("ssh remote url with dot should return", func() {
			org, repo := GetOrganizationAndRepository("git@bitbucket.com:Praqma/praqma.com.git")

			Ω(org).Should(Equal("Praqma"))
			Ω(repo).Should(Equal("praqma.com"))
		})

	})

	ginkgo.Context("protocol tests", func() {

		ginkgo.It("Testing with expected result set git-phlow", func() {
			org, repo := GetOrganizationAndRepository("git@github.com:Praqma/git-phlow.git")

			Ω(org).Should(Equal("Praqma"))
			Ω(repo).Should(Equal("git-phlow"))
		})

		ginkgo.It("Testing with expected result set grit", func() {
			org, repo := GetOrganizationAndRepository("https://github.com/cho45/grit.git")

			Ω(org).Should(Equal("cho45"))
			Ω(repo).Should(Equal("grit"))
		})

		ginkgo.It("Testing with expected result set hops", func() {
			org, repo := GetOrganizationAndRepository("git://github.com/koke/hops.git")

			Ω(org).Should(Equal("koke"))
			Ω(repo).Should(Equal("hops"))
		})

		ginkgo.It("Testing with expected result set helmsman", func() {
			org, repo := GetOrganizationAndRepository("ssh://git@bitbucket:7560/praqma/helmsman.git")

			Ω(org).Should(Equal("praqma"))
			Ω(repo).Should(Equal("helmsman"))
		})

	})

})
