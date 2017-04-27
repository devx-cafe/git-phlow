package githandler

import (
	"io/ioutil"
	"testing"

	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRemote(t *testing.T) {
	SkipConvey("Running tests on 'Remote' function (runs in project)", t, func() {

		Convey("Remote should return organisation and repo name", func() {
			remote, _ := Remote()

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

func TestRemoteUrlExtractor(t *testing.T) {
	Convey("Running test on 'remoteUrlExtractor' function", t, func() {
		Convey("https remote url should return", func() {
			var https = "https://gh.com/Org/sOme--repo.git"
			info := remoteURLExtractor(https)
			So(info.Organisation, ShouldEqual, "Org")
			So(info.Repository, ShouldEqual, "sOme--repo")

		})

		Convey("ssh remote url should return", func() {
			var ssh = "git@gh.com:Org/some-repo.git"
			info := remoteURLExtractor(ssh)
			So(info.Organisation, ShouldEqual, "Org")
			So(info.Repository, ShouldEqual, "some-repo")
		})

		Convey("ssh remote url with dot should return", func() {
			var ssh = "git@gh.com:Praqma/praqma.com.git"
			info := remoteURLExtractor(ssh)
			t.Log(info.Organisation)
			t.Log(info.Repository)
			So(info.Organisation, ShouldEqual, "Praqma")
			So(info.Repository, ShouldEqual, "praqma.com")
		})

	})
}

func TestStatusPorcelain(t *testing.T) {

	SkipConvey("Running test on 'StatusPorcelain' function", t, func() {

		Convey("Status should return one of behind and forward", func() {

			out, err := StatusPorcelain()
			t.Log(out)
			t.Log(err)
			So(out, ShouldEqual, "## 52-implement-a-park-feature...origin/master [ahead 1]")
		})
	})
}
