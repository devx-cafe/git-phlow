package plugins

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRepoNameAndOrg(t *testing.T) {
	Convey("Test Function GetRepoNameAndOrg", t, func() {

		Convey("Organization/name and Repo Name should be extracted", func() {
			var fetchUrl = "origin	git@github.com:Praqma/git-phlow.git (fetch)"

			name, org := GetRepoAndUser(fetchUrl)

			So(org, ShouldEqual, "Praqma")
			So(name, ShouldEqual, "git-phlow")
		})
	})
}
