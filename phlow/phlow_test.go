package phlow

import (
	"testing"

	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/praqma/git-phlow/githandler"
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

func TestUpNext(t *testing.T) {
	Convey("Running tests on 'GetNextBranch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of clean function", func() {
			branches := githandler.BranchReady("origin")
			res := GetNextBranch(branches)

			So(res, ShouldEqual, "origin/ready/15-issue-branch")
		})

		testfixture.RemoveTestRepository(t)

	})
}
