package githandler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemote(t *testing.T) {
	Convey("Running tests on 'Remote' function", t, func() {

		Convey("Remote should return organisation and repo name", func() {
			remote, err := Remote()

			So(err, ShouldBeNil)
			So(remote.Repository, ShouldEqual, "git-phlow")
			So(remote.Organisation, ShouldEqual, "Praqma")
		})
	})
}
