package phlow

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIsBehindOrAhead(t *testing.T) {
	Convey("Running tests on 'isBehind' function", t, func() {

		Convey("behind should return behind=true, ahead=false", func() {
			in := "## 52-implement-a-park-feature...origin/master [behind 13]"

			bh, ah := isBehindOrAhead(in)
			So(bh, ShouldBeTrue)
			So(ah, ShouldBeFalse)
		})

		Convey("no behind should behind=false, ahead=false", func() {
			in := "## 52-implement-a-park-feature...origin/master"

			bh, ah := isBehindOrAhead(in)
			So(bh, ShouldBeFalse)
			So(ah, ShouldBeFalse)
		})

		Convey("ahead should return behind=false, ahead=true", func() {
			in := "## 52-implement-a-park-feature...origin/master [ahead 1]"

			bh, ah := isBehindOrAhead(in)
			So(bh, ShouldBeFalse)
			So(ah, ShouldBeTrue)
		})
		Convey("no ahead should return behind=false, ahead=false", func() {
			in := "## 52-implement-a-park-feature...origin/master"

			bh, ah := isBehindOrAhead(in)
			So(bh, ShouldBeFalse)
			So(ah, ShouldBeFalse)
		})

		Convey("behind and ahead should return behind=true, ahead=true", func() {
			in := "## delivered/2-build-a-deployment-platform...origin/master [ahead 1, behind 1]"

			bh, ah := isBehindOrAhead(in)
			So(bh, ShouldBeTrue)
			So(ah, ShouldBeTrue)
		})
	})
}
