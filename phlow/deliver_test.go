package phlow

import (
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTestDeliver(t *testing.T) {
	Convey("Running tests on 'TestDeliver' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Wrong path to script should return error", func() {
			err := TestDeliver([]string{"./lalalal"})
			t.Log(err)
			So(err, ShouldNotBeNil)
		})

		Convey("Right path to script should not return error", func() {
			err := TestDeliver([]string{"./test.sh"})
			t.Log(err)
			So(err, ShouldBeNil)
		})

		Convey("Right path to error script should return error", func() {
			err := TestDeliver([]string{"./testerr.sh"})
			t.Log(err)
			So(err, ShouldNotBeNil)
		})

		Convey("Valid one line command should not return error", func() {
			options.GlobalFlagShowTestOutput = true
			err := TestDeliver([]string{"ls"})
			t.Log(err)
			So(err, ShouldBeNil)
		})

		Convey("valid two line command should not return error", func() {
			options.GlobalFlagShowTestOutput = true
			err := TestDeliver([]string{"ls", "-lah"})
			t.Log(err)
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)

	})
}

func TestConvertCommand(t *testing.T) {
	Convey("Running tests on 'ConvertCommand' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("input 'path' should return path only", func() {
			cmd, args := convertCommand([]string{"./path/to/script.sh"})

			So(cmd, ShouldEqual, "./path/to/script.sh")
			So(args, ShouldBeEmpty)
		})

		Convey("multi argument should return command and arguments", func() {
			cmd, args := convertCommand([]string{"multi", "line", "command", "and", "args"})

			So(cmd, ShouldEqual, "multi")
			So(args, ShouldContain, "line")
			So(args, ShouldContain, "command")
			So(args, ShouldContain, "and")
			So(args, ShouldContain, "args")
		})

		testfixture.RemoveTestRepository(t)
	})
}
