package executor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRunCommand(t *testing.T) {

	Convey("Runnig tests on 'ExecuteCommand' function", t, func() {

		Convey("running ls should not return an error and stdout", func() {
			output, err := ExecuteCommand("ls", "-lah")
			So(output, ShouldNotBeBlank)
			So(err, ShouldBeNil)
		})

		Convey("running lsk should return an error and stderr", func() {
			output, err := ExecuteCommand("lsk", "-lah")
			So(output, ShouldBeBlank)
			So(err, ShouldNotBeNil)
		})
	})
}
