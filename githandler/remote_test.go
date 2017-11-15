package githandler

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Remote", func() {

	Describe("orgAndRepo", func() {

		Context("using github", func() {
			It("should return organisation and repository with https url", func() {
				var https = "https://github.com/Praqma/phlow-test-strategy.git"
				info := OrgAndRepo(https)
				Ω(info.Organisation).Should(Equal("Praqma"))
				Ω(info.Repository).Should(Equal("phlow-test-strategy"))
			})

			It("should return organisation and repository with ssh", func() {
				var ssh = "git@github.com:Org/some-repo.git"
				info := OrgAndRepo(ssh)
				Ω(info.Organisation).Should(Equal("Org"))
				Ω(info.Repository).Should(Equal("some-repo"))
			})

			It("ssh remote url with dot should return", func() {
				var ssh = "git@github.com:Praqma/praqma.com.git"
				info := OrgAndRepo(ssh)
				Ω(info.Organisation).Should(Equal("Praqma"))
				Ω(info.Repository).Should(Equal("praqma.com"))
			})
		})

		Context("using bitbucket", func() {

			It("should return organisation and repository with https url", func() {
				var https = "https://bitbucket.com/Org/sOme--repo.git"
				info := OrgAndRepo(https)
				Ω(info.Organisation).Should(Equal("Org"))
				Ω(info.Repository).Should(Equal("sOme--repo"))
			})

			It("should return organisation and repository with ssh", func() {
				var ssh = "git@bitbucket.com:Org/some-repo.git"
				info := OrgAndRepo(ssh)
				Ω(info.Organisation).Should(Equal("Org"))
				Ω(info.Repository).Should(Equal("some-repo"))
			})

			It("ssh remote url with dot should return", func() {
				var ssh = "git@bitbucket.com:Praqma/praqma.com.git"
				info := OrgAndRepo(ssh)
				Ω(info.Organisation).Should(Equal("Praqma"))
				Ω(info.Repository).Should(Equal("praqma.com"))
			})

		})

		Context("protocol tests", func() {

			type Case struct {
				Expected string
				Result   string
			}

			It("Testing with expected result set git-phlow", func() {
				c := Case{"git@github.com:Praqma/git-phlow.git", "Praqma:git-phlow"}
				Ω(OrgAndRepo(c.Expected).Organisation).Should(Equal(strings.Split(c.Result, ":")[0]))
				Ω(OrgAndRepo(c.Expected).Repository).Should(Equal(strings.Split(c.Result, ":")[1]))

			})
			It("Testing with expected result set grit", func() {
				c := Case{"https://github.com/cho45/grit.git", "cho45:grit"}
				Ω(OrgAndRepo(c.Expected).Organisation).Should(Equal(strings.Split(c.Result, ":")[0]))
				Ω(OrgAndRepo(c.Expected).Repository).Should(Equal(strings.Split(c.Result, ":")[1]))

			})
			It("Testing with expected result set hops", func() {
				c := Case{"git://github.com/koke/hops.git", "koke:hops"}
				Ω(OrgAndRepo(c.Expected).Organisation).Should(Equal(strings.Split(c.Result, ":")[0]))
				Ω(OrgAndRepo(c.Expected).Repository).Should(Equal(strings.Split(c.Result, ":")[1]))

			})
			It("Testing with expected result set helmsman", func() {
				c := Case{"ssh://git@bitbucket:7560/praqma/helmsman.git", "praqma:helmsman"}
				Ω(OrgAndRepo(c.Expected).Organisation).Should(Equal(strings.Split(c.Result, ":")[0]))
				Ω(OrgAndRepo(c.Expected).Repository).Should(Equal(strings.Split(c.Result, ":")[1]))

			})

		})

	})

})
