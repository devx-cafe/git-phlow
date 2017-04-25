package executor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"bytes"
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

func TestExecPipeCommand(t *testing.T) {

	Convey("Runnig tests on 'ExecPipeCommand' function", t, func() {

		SkipConvey("should run with any number of commands", func() {

			var buf bytes.Buffer
			err := ExecPipeCommand(&buf,
				exec.Command("ls", "-lah"),
				exec.Command("grep", "c"),
				exec.Command("sort", "-r"))

			So(err, ShouldBeNil)
			So(buf.String(), ShouldNotBeEmpty)
		})

		SkipConvey("should run with two commands", func() {
			var buf bytes.Buffer

			err := ExecPipeCommand(&buf,
				exec.Command("ls", "-lah"),
				exec.Command("grep", "c"))

			So(err, ShouldBeNil)
			So(buf.String(), ShouldNotBeEmpty)
		})

		SkipConvey("should run with one command", func() {
			var buf bytes.Buffer

			err := ExecPipeCommand(&buf, exec.Command("ls", "-lah"))

			So(err, ShouldBeNil)
			So(buf.String(), ShouldNotBeEmpty)
		})

		Convey("First function should error", func() {
			var buf bytes.Buffer

			err := ExecPipeCommand(&buf,
				exec.Command("argh", "blash"),
				exec.Command("grep", "stuff"))

			So(err, ShouldNotBeNil)
			So(buf.String(), ShouldBeEmpty)
		})

		Convey("Second function should error", func() {
			var buf bytes.Buffer

			err := ExecPipeCommand(&buf,
				exec.Command("ls", "-lah"),
				exec.Command("jklasd", "stuff"))

			So(err, ShouldNotBeNil)
			So(buf.String(), ShouldBeEmpty)
		})

	})
}
