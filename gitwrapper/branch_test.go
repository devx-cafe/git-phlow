/*
Test package for gitwrapper Branch
*/
package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"strings"
	"testing"

	"github.com/praqma/git-phlow/testfixture"
)

func TestBranch(t *testing.T) {
	Convey("Test ListBranch", t, func() {

		testfixture.SetupTestRepo()

		Convey("Listbranch should return master", func() {

			branch, err := InitGit().Branch().ListBranches()
			master := false

			for _, br := range branch {
				t.Log(br)
				if strings.Contains(br, "master") {
					master = true
				}
			}
			So(master, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}

func TestCreateBranch(t *testing.T) {
	Convey("Test creation of branch", t, func() {

		testfixture.SetupTestRepo()

		Convey("Create branch testphlow", func() {
			accessBranch := InitGit().Branch()
			branch, err := accessBranch.CreateBranch("testphlow")
			list, errList := accessBranch.ListBranches()

			newBranch := false
			for _, br := range list {
				if strings.Contains(br, branch) {
					newBranch = true
				}
			}

			So(newBranch, ShouldBeTrue)
			So(err, ShouldBeNil)
			So(errList, ShouldBeNil)
		})

		Convey("Fail to create already existing branch bar", func() {
			foofail, err := InitGit().Branch().CreateBranch("bar")
			So(foofail, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}

func TestCurrentBranch(t *testing.T) {

	Convey("Test GetCurrent branch", t, func() {

		testfixture.SetupTestRepo()

		Convey("current branch should be master", func() {
			brancher := InitGit().Branch()
			branch, err := brancher.CurrentBranch()

			So(branch, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}
