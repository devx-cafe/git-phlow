/*
Test package for gitwrapper Branch
*/
package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
	"strings"
	"github.com/praqma/git-phlow/testfixture"
	"fmt"
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

			if err != nil {
				fmt.Println(err)
			}

			var master = false
			for _, br := range branch {
				if strings.Contains(br, "master") {
					master = true
				}
			}
			So(master, ShouldBeTrue)
		})
		
		textfixture.TearDownTestRepo()
	})
}
