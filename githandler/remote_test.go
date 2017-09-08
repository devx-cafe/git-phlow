package githandler

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	})

})
