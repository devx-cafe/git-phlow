package executor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"bytes"
	"os/exec"
	"io"
	"os"
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

func TestExecutePipe(t *testing.T) {
	Convey("Running tests on 'ExecutePipe', function ", t, func() {
		Convey("piping commands should not return error", func() {

			var output bytes.Buffer

			err := ExecutePipe(&output,
				exec.Command("ls", "-lah"),
				exec.Command("grep","."),
				exec.Command("sort", "-r"),
			)

			io.Copy(os.Stdout, &output)

			t.Log(err)
		})
	})
}
