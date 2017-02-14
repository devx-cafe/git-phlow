package phlowimpl

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestGetIssues(t *testing.T) {

	Convey("Test Get issue mappings", t, func() {

		Convey("map should contain issue branch", func() {

			branches := []string{"1-a-simple-branch", "22-feature-is-cool?"}
			var branchMap = make(map[int]string)

			getIssues(branches, branchMap)

			So(branches[0], ShouldEqual, branchMap[1])
			So(branches[1], ShouldEqual, branchMap[22])
			So(branchMap[3], ShouldBeBlank)
		})
	})
}
