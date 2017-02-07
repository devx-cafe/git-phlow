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



		Convey("run: 'git sts' - should return ExitCode, stderr and err", func() {


		})
	})

}

func TestIsInPath(t *testing.T) {

	Convey("Test function IsInPath", t, func() {

		var cd string = "cd"                    //Unix, Darwin, windows should all have 'cd'
		var notAnApp string = "libblobdibdab" //Random string which s unlikely to be an app


		Convey("Test app" + cd + "is in path", func() {
			actual := IsInPath(cd)
			So(actual, ShouldBeNil)
		})

		Convey("Test program " + notAnApp + "is not in path", func() {

			actual := IsInPath(notAnApp)
			So(actual, ShouldNotBeNil)
		})

	})
}
