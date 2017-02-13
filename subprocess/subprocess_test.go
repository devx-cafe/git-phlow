package subprocess

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleExec(t *testing.T) {

	Convey("Test function SimpleExec", t, func() {

		Convey("run: 'ls' - should return no errors ", func() {

			output, err := SimpleExec("ls", "-lah")

			So(output, ShouldNotBeBlank)
			So(err, ShouldBeNil)
		})

	})

}

func TestIsInPath(t *testing.T) {

	Convey("Test function IsInPath", t, func() {

		var ls string = "ls"                    //Unix, Darwin, windows should all have 'ls'
		var notAnApp string = "libblobdibdab" //Random string which s unlikely to be an app


		Convey("Test app" + ls + "is in path", func() {
			actual := IsInPath(ls)
			So(actual, ShouldBeNil)
		})

		Convey("Test program " + notAnApp + "is not in path", func() {

			actual := IsInPath(notAnApp)
			So(actual, ShouldNotBeNil)
		})

	})
}
