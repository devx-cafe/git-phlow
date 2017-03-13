package phlow

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/praqma/git-phlow/testfixture"
	"github.com/praqma/git-phlow/options"
)

func TestGetIssueFromBranch(t *testing.T) {
	Convey("Running tests on 'GetIssuesFromBranch'", t, func() {

		Convey("GetIssueSFromBranch should return 1", func() {

			i := GetIssueFromBranch("39-enable---Add-sign-in-function")
			So(i, ShouldEqual, 39)
		})
	})
}

func TestClean(t *testing.T) {

	Convey("Runnign tests on 'Clean' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of local clean function", func() {
			options.GlobalFlagLocal = true
			Clean("origin")

		})

		testfixture.RemoveTestRepository(t)
	})

}

func TestCleanRemote(t *testing.T) {

	Convey("Runnign tests on 'Clean' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of clean function", func() {
			options.GlobalFlagLocal = false
			Clean("origin")

		})

		testfixture.RemoveTestRepository(t)
	})

}
