package githandler

import (
	"io/ioutil"
	"testing"

	"github.com/praqma/git-phlow/testfixture"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemote(t *testing.T) {
	Convey("Running tests on 'Remote' function (runs in project)", t, func() {

		Convey("Remote should return organisation and repo name", func() {
			remote, _ := Remote("master")

			So(remote.Repository, ShouldEqual, "git-phlow")
			So(remote.Organisation, ShouldEqual, "Praqma")
		})

	})
}

func TestConfig(t *testing.T) {
	Convey("Running tests on 'ConfigBranchRemote'", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("ConfigBranchRemote should return origin", func() {
			remote := ConfigBranchRemote("master")

			So(remote, ShouldEqual, "origin")
		})

		Convey("ConfigBranchRemote of wrong branch should return err", func() {
			remote := ConfigBranchRemote("bsld")

			So(remote, ShouldEqual, "")
		})

		testfixture.RemoveTestRepository(t)
	})
}

func TestBranch(t *testing.T) {
	Convey("Running tests on 'Branch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("branch should return List of branches", func() {
			info, err := Branch()
			So(len(info.List), ShouldEqual, 9)
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
			So(info.List, ShouldHaveLength, 8)
		})

		Convey("BranchDelete should delete remote branch and return message", func() {
			_, err1 := BranchDelete("delivered/24-issue-branch", "origin", true, false)
			_, err2 := BranchDelete("delivered/42-issue-branch", "origin", true, false)
			info, _ := Branch()

			t.Log(info.List)
			So(err1, ShouldBeNil)
			So(err2, ShouldBeNil)
			So(info.List, ShouldHaveLength, 6)
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

func TestCheckout(t *testing.T) {

	Convey("Runnign tests in 'Checkout' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Checkout existing branch should not return error", func() {
			err := CheckOut("bar")
			So(err, ShouldBeNil)
		})

		Convey("Checkout Current branch should not return error", func() {
			info, _ := Branch()
			err := CheckOut(info.Current)
			So(err, ShouldBeNil)
		})

		Convey("Checkout from origin branch should not return error", func() {
			err := CheckOut("foo")
			So(err, ShouldBeNil)
		})

		Convey("Checkout nonexisting branch should return error", func() {
			err := CheckOut("i-am-not-a-branch")
			So(err, ShouldNotBeNil)
		})

		Convey("Checkout uncomitted changes should return error", func() {
			ioutil.WriteFile("./README.md", []byte("I AM A CONFLICTIONG CHANGE"), 0755)
			err := CheckOut("foo")
			So(err, ShouldNotBeNil)
		})

		Convey("Checkout new origin branch should not return error", func() {
			err := CheckoutNewBranchFromRemote("12-issue", "master")
			So(err, ShouldBeNil)
		})

		Convey("Checkout existing origin branch should return error", func() {
			err := CheckoutNewBranchFromRemote("foo", "master")
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)
	})
}

func TestFetch(t *testing.T) {

	Convey("Runnig tests in 'Fetch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Fetch all should not return error", func() {
			err := Fetch()
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)
	})
}

func TestStatus(t *testing.T) {

	Convey("Running test on 'Status' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Status should not return error", func() {
			err := Status()
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)
	})
}
