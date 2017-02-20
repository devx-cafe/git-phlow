package githandler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/praqma/git-phlow/testfixture"
	"strings"
	"io/ioutil"
	"gopkg.in/libgit2/git2go.v25"
)

func TestRemote(t *testing.T) {
	Convey("Running tests on 'Remote' function", t, func() {

		Convey("Remote should return organisation and repo name", func() {
			remote, err := Remote()
			So(err, ShouldBeNil)
			So(remote.Repository, ShouldEqual, "git-phlow")
			So(remote.Organisation, ShouldEqual, "Praqma")
		})
	})
}

func TestBranch(t *testing.T) {
	Convey("Running tests on 'Branch' function", t, func() {

		testfixture.CreateTestRepository(t,true)

		Convey("branch should return list of branches", func() {
			info, err := Branch("list")
			So(len(info.list), ShouldEqual, 3)
			So(err, ShouldBeNil)
		})

		Convey("branch should return current branch", func() {
			info, err := Branch("current")
			So(info.current, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})

	})
}

func TestCheckout(t *testing.T) {

	Convey("Runnign tests in 'Checkout' function", t, func() {

		Convey("Checkout existing branch should not return error", func() {
			err := CheckOut("bar", false)
			So(err, ShouldBeNil)
		})

		Convey("Checkout current branch should not return error", func() {
			info, _ := Branch("current")
			err := CheckOut(info.current, false)
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

	})
}

func TestFetch(t *testing.T) {

	Convey("Runnig tests in 'Fetch' function", t, func() {

		Convey("Fetch all should not return error", func() {
			err := Fetch()
			So(err, ShouldBeNil)
		})
	})
}

func TestStatus(t *testing.T) {

	Convey("Running test on 'Status' function", t, func() {

		Convey("Status should not return error", func() {
			err := Status()
			So(err, ShouldBeNil)
		})
	})
}
