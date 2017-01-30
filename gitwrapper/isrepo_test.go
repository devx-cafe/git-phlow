package gitwrapper

import (
	"testing"

	 testConstant "github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"

)


//Tests rely on $GOPATH, it uses gopath to find the project directory and
//having a place where the fixture can create a test folder
func TestGetCurrentDirectory(t *testing.T) {

	Convey("Test method GetCurrentDirectory",t, func() {

		Convey("Directory Should contain root folder", func() {
			path, err := GetCurrentDirectory()

			So(err, ShouldBeNil)
			So(path, ShouldContainSubstring, "git-phlow")
		})
	})
}

func TestIsRepository(t *testing.T) {

	Convey("Test method IsRepoInitialized on project directory", t, func() {

		Convey("git-phlow should be initialized", func() {

			var path string
			var isRepo bool

			path = testConstant.ProjectDirectory
			isRepo = IsRepository(path)

			So(isRepo, ShouldBeTrue)
		})
	})
}
