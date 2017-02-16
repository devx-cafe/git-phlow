package phlow

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
	"github.com/praqma/git-phlow/testfixture"
	"github.com/praqma/git-phlow/gitwrapper"
)

func TestGetIssues(t *testing.T) {

	Convey("Test Get issue mappings", t, func() {

		Convey("map should contain issue branch", func() {

			branches := []string{"1-a-simple-branch", "22-feature-is-cool?"}
			branchMap := getBranchesAsMap(branches)

			So(branches[0], ShouldEqual, branchMap[1])
			So(branches[1], ShouldEqual, branchMap[22])
			So(branchMap[3], ShouldBeBlank)
		})
	})
}

func TestSwitchOrReworkExistingBranch(t *testing.T) {

	Convey("Test SwitchOrReworkBranch function", t, func() {

		testfixture.SetupTestRepo()

		Convey("Checkout branch from exisitng issue", func() {

			var branchFromFixture = "11-issue-bar"
			git := gitwrapper.InitGit()
			err := SwitchOrReworkExistingBranch(branchFromFixture, git)

			t.Log(err)
			So(err, ShouldBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}
