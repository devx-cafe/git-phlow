/*
Test package for gitwrapper Branch
*/
package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
	"strings"
	"github.com/praqma/git-phlow/testfixture"
)

func TestStringConcat(t *testing.T) {

	Convey("Test function effecientStringConcat", t, func() {

		Convey("Test variable inputs get concatenated correctly", func() {
			var expectedLong = "created 'something' new"
			var actualLong = efficientConcatString("created '", "something", "' new")

			So(actualLong, ShouldEqual, expectedLong)

		})

		Convey("Test funny signs gets concatenated as well", func() {

			var expectedShort = "j$¢‰¿≈¯¯¯"
			var actualShort = efficientConcatString("j$¢‰¿≈", "¯¯¯")

			So(expectedShort, ShouldEqual, actualShort)

		})
	})
}

func TestBranch(t *testing.T) {
	Convey("Test function NewBranch and Branch", t, func() {

		textfixture.SetupTestRepo()

		Convey("Test function 'Branch' should contain master branch", func() {

			branch, err := InitGit().Branch().ListBranches()

			master, foo := false, false

			for _, br := range branch {
				if strings.Contains(br, "master") {
					master = true
				}
				if strings.Contains(br, "foo") {
					foo = true
				}
			}
			So(master, ShouldBeTrue)
			So(foo, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		textfixture.TearDownTestRepo()
	})
}

func CreateBranch(t *testing.T) {
	Convey("Test creation of branch", t, func() {
		textfixture.SetupTestRepo()
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
		textfixture.TearDownTestRepo()
	})

}
