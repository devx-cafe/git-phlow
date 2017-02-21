package githandler

import (
	"testing"
	"github.com/praqma/git-phlow/testfixture"
	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemote(t *testing.T) {
	Convey("Running tests on 'Remote' function (runs in project)", t, func() {

		Convey("Remote should return organisation and repo name", func() {
			remote, err := Remote()
			So(err, ShouldBeNil)

			t.Log(err)


			So(remote.Repository, ShouldEqual, "git-phlow")
			So(remote.Organisation, ShouldEqual, "Praqma")
		})
	})
}

func TestBranch(t *testing.T) {
	Convey("Running tests on 'Branch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("branch should return List of branches", func() {
			info, err := Branch("list")
			So(len(info.List), ShouldEqual, 2)
			So(err, ShouldBeNil)
		})

		Convey("branch should return Current branch", func() {
			info, err := Branch("current")
			So(info.Current, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)

	})
}

func TestCheckout(t *testing.T) {

	Convey("Runnign tests in 'Checkout' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Checkout existing branch should not return error", func() {
			err := CheckOut("bar", false)
			So(err, ShouldBeNil)
		})

		Convey("Checkout Current branch should not return error", func() {
			info, _ := Branch("Current")
			err := CheckOut(info.Current, false)
			So(err, ShouldBeNil)
		})

		Convey("Checkout from origin branch should not return error", func() {
			err := CheckOut("foo", false)
			So(err, ShouldBeNil)
		})

		Convey("Checkout nonexisting branch should return error", func() {
			err := CheckOut("i-am-not-a-branch", false)
			So(err, ShouldNotBeNil)
		})

		Convey("Checkout uncomitted changes should return error", func() {
			ioutil.WriteFile("./README.md", []byte("I AM A CONFLICTIONG CHANGE"), 0755)
			err := CheckOut("foo", false)
			So(err, ShouldNotBeNil)
		})

		Convey("Checkout now origin branch should not return error", func() {
			err := CheckOut("12-issue", true)
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
