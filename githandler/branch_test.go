package githandler

import (
	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBranch(t *testing.T) {
	Convey("Running tests on 'Branch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("branch should return List of branches", func() {
			info, err := Branch()
			So(len(info.List), ShouldEqual, 11)
			So(err, ShouldBeNil)
		})

		Convey("branch should return Current branch", func() {
			info, err := Branch()
			So(info.Current, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)

	})
}

func TestBranchDelete(t *testing.T) {

	testfixture.CreateTestRepository(t, false)

	Convey("Running tests on 'BranchDelete' function", t, func() {

		Convey("BranchDelte should delete local branch and return message", func() {
			output, err := BranchDelete("delivered/1-issue-branch", "", false, false)
			info, _ := Branch()

			t.Log(info.List)
			So(err, ShouldBeNil)
			So(output, ShouldNotBeEmpty)
			So(info.List, ShouldHaveLength, 10)
		})

		Convey("BranchDelete should delete remote branch and return message", func() {
			_, err1 := BranchDelete("delivered/24-issue-branch", "origin", true, false)
			_, err2 := BranchDelete("delivered/42-issue-branch", "origin", true, false)
			info, _ := Branch()

			t.Log(info.List)
			So(err1, ShouldBeNil)
			So(err2, ShouldBeNil)
			So(info.List, ShouldHaveLength, 8)
		})
	})

	testfixture.RemoveTestRepository(t)
}

func TestBranchDelivered(t *testing.T) {

	testfixture.CreateTestRepository(t, false)

	Convey("Running tests in 'BranchDelivered' function", t, func() {
		locals, remotes := BranchDelivered("origin")
		So(locals, ShouldHaveLength, 1)
		So(remotes, ShouldHaveLength, 2)
		So(remotes, ShouldContain, "delivered/24-issue-branch")
	})

	testfixture.RemoveTestRepository(t)
}

func TestBranchReady(t *testing.T) {

	testfixture.CreateTestRepository(t, false)

	Convey("Running tests in 'BranchDelivered' function", t, func() {
		remotes := BranchReady("origin")
		t.Log(remotes)
		So(remotes, ShouldHaveLength, 2)
		So(remotes, ShouldContain, "origin/ready/99-issue-branch")
	})

	testfixture.RemoveTestRepository(t)
}

func TestBranchTime(t *testing.T) {
	Convey("Running tests on 'BranchTime' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Should get unix timestamp", func() {
			output, err := BranchTime("origin/ready/99-issue-branch")

			So(err, ShouldBeNil)
			So(output, ShouldBeGreaterThan, 100000)
		})

		Convey("Should fail geting unix timestamp", func() {
			output, err := BranchTime("bluarh.. not a branch")

			So(err, ShouldNotBeNil)
			So(output, ShouldEqual, -1)
		})

		testfixture.RemoveTestRepository(t)

	})
}
