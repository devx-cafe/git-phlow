package phlow

import (
	"testing"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIssues(t *testing.T) {
	Convey("Running tests on 'Issues' command", t, func() {
		Convey("testing output", func() {
			IssueList()

		})
	})
}

func TestUpNext(t *testing.T) {
	Convey("Running tests on 'GetNextBranch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of clean function", func() {
			branches := githandler.BranchReady("origin")
			res := getNextBranch(branches)
			So(res, ShouldEqual, "ready/15-issue-branch")
		})

		testfixture.RemoveTestRepository(t)

	})
}

func TestRemoveRemoteFromName(t *testing.T) {
	Convey("Running tests on 'remoteRemoteFromName' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("origin should be removed", func() {
			branches := githandler.BranchReady("origin")
			res := getNextBranch(branches)
			res = removeRemoteFromUpNext(res)

			t.Log(res)

			So(res, ShouldEqual, "ready/15-issue-branch")
		})

		testfixture.RemoveTestRepository(t)

	})
}
