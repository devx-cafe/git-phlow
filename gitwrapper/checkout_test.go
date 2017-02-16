package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"io/ioutil"
	"testing"

	"github.com/praqma/git-phlow/testfixture"
)

func TestCheckout(t *testing.T) {

	Convey("Checkout Test", t, func() {

		testfixture.SetupTestRepo()

		Convey("Checkout branch bar should return branch name on success", func() {
			str, err := InitGit().Checkout().Checkout("bar")

			So(str, ShouldEqual, "bar")
			So(err, ShouldBeNil)
		})

		Convey("Checkout current branch should not return error", func() {
			git := InitGit()
			current, _ := git.Branch().CurrentBranch()
			_, err := git.Checkout().Checkout(current)

			So(err, ShouldBeNil)
		})

		Convey("checkout from origin branch should be return message", func() {
			str, err := InitGit().Checkout().Checkout("foo")
			So(str, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("Checkout nonexisting branch should fail", func() {
			git := InitGit()
			_, err := git.Checkout().Checkout("non-existing-branch")

			So(err, ShouldNotBeNil)
		})

		Convey("Checkout whith uncomitted changes", func() {
			git := InitGit()
			ioutil.WriteFile("./README.md", []byte("I AM A CONFLICTIONG CHANGE"), 0755)

			_, err := git.Checkout().Checkout("foo")

			So(err, ShouldNotBeNil)
		})

		Convey("Checkout new branch should from origin should fail", func() {
			git := InitGit()
			output, err := git.Checkout().CheckoutNewBranchFromOrigin("some-branch-name", "origin/master")
			t.Log(err)
			t.Log(output)
			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()

	})
}
