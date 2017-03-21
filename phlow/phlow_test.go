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

			So(res, ShouldEqual, "origin/ready/15-issue-branch")
		})

		testfixture.RemoveTestRepository(t)

	})
}
