package githandler

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Branch", func() {

	Describe("AsList", func() {

		It("should return a list of branches", func() {
			info := AsList(gitBranchOut)
			Ω(info.List).Should(HaveLen(11))
		})

		It("should return nothing", func() {
			info := AsList("læadsjklfjaskldjsakldjaskld")
			Ω(info.List).Should(HaveLen(1))
		})

		It("should return the current branch ", func() {
			info := AsList(gitBranchOut)
			Ω(info.Current).Should(Equal("master"))
		})

		It("should not return master", func() {
			info := AsList("læadsjklfjaskldjsakldjaskld")
			Ω(info.Current).Should(Equal(""))
		})

	})

	Describe("Delivered", func() {

		It("should return delivered branch", func() {
			info := AsList(gitBranchOut)
			local, remote := Delivered(info, "origin")
			Ω(local).Should(HaveLen(2))
			Ω(remote).Should(HaveLen(1))
		})

	})

	Describe("Ready", func() {

		It("should return list of ready branches", func() {
			info := AsList(gitBranchOut)
			ready := Ready(info, "origin", "ready")

			Ω(ready).Should(HaveLen(1))
			Ω(ready).Should(ContainElement("origin/ready/31-add-documentation-for-docker-compose"))
		})

	})

})

var gitBranchOut = `
  23-add-impl-for-postgres
  24-isolate-rest-for-scalability
  31-add-documentation-for-docker-compose
  delivered/30-change-docker-config-to-k8
  delivered/31-add-documentation-for-docker-compose
* master
  remotes/origin/HEAD -> origin/master
  remotes/origin/master
  remotes/origin/ready/31-add-documentation-for-docker-compose
  remotes/origin/version
  remotes/origin/delivered/31-add-documentation-for-docker-compose
`