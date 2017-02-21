package phlow

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetIssueFromBranch(t *testing.T) {
	Convey("Running tests on 'GetIssuesFromBranch'", t, func() {

		Convey("GetIssueSFromBranch should return 1", func() {

			i := GetIssueFromBranch("39-enable---Add-sign-in-function")
			So(i, ShouldEqual, 39)
		})
	})
}
