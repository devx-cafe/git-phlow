package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"github.com/praqma/git-phlow/testfixture"

	"testing"
)

func TestCheckout(t *testing.T) {

	Convey("Checkout Test", t, func() {

		testfixture.SetupTestRepo()

		Convey("Checkout other branch should be possible", func() {
			str, err := InitGit().Checkout().Checkout("foo")

			So(str, ShouldEqual, "foo")
			So(err, ShouldBeNil)
		})

		Convey("Checkout current branch should not result in error", func() {
			git := InitGit()
			current, _ := git.Branch().CurrentBranch()
			checkout, err := git.Checkout().Checkout(current)

			So(err, ShouldBeNil)
			So(checkout, ShouldEqual, current)

		})

		Convey("Checkout nonexisting branch should fail", func() {
			git := InitGit()
			_, err := git.Checkout().Checkout("non-existing-branch")

			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()

	})
}
