package gitwrapper

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

)

func TestStatus_Status(t *testing.T) {

	Convey("Test Status methods", t, func() {

		Convey("Status should not return error", func() {
			err := NewStatus().Status()
			So(err, ShouldBeNil)
		})
	})
}
