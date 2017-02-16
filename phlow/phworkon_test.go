package phlow

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
	"github.com/praqma/git-phlow/testfixture"
	"github.com/praqma/git-phlow/gitwrapper"
	"github.com/praqma/git-phlow/plugins"
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

		Convey("Cheking out existing branch should resume work", func() {

			var branchFromFixture = "11-issue-bar"
			git := gitwrapper.InitGit()
			err := SwitchOrReworkExistingBranch(branchFromFixture, git)

			t.Log(err)
			So(err, ShouldBeNil)
		})

		Convey("Checking out non-issue branch should not resume work", func() {

			var branchFromFixture = "looo"
			git := gitwrapper.InitGit()
			err := SwitchOrReworkExistingBranch(branchFromFixture, git)

			t.Log(err)
			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}

func TestCheckoutNewBranchFromPluginIssue(t *testing.T) {

	Convey("Test Checkout on a branch created from an issue", t, func() {

		testfixture.SetupTestRepo()

		Convey("Checkout with valid issue number should checkout new branch", func() {
			plugin := NewRepository()
			git := gitwrapper.InitGit()

			var issueFromGithub = 12

			err := CheckoutNewBranchFromPluginIssue(issueFromGithub, plugin, git)

			t.Log(err)

			So(err, ShouldBeNil)
		})

		Convey("Checkout with non-existing should return error", func() {
			plugin := NewRepository()
			git := gitwrapper.InitGit()

			err := CheckoutNewBranchFromPluginIssue(11, plugin, git)

			t.Log(err)

			So(err, ShouldNotBeNil)
		})

		testfixture.TearDownTestRepo()
	})
}

type FakePlugin struct {
}

func NewRepository() plugins.Plugin {
	return &FakePlugin{}
}

func (f *FakePlugin) DefaultBranch() string {
	return "master"
}
func (f *FakePlugin) ListIssues() map[int]string {
	maps := make(map[int]string)
	maps[12] = "I am an issue"
	return maps
}

func (f *FakePlugin) SetAssignee(name string) {

}
func (f *FakePlugin) SetLabelsOnIssue(args ...string) {

}
func (f *FakePlugin) InitializeRepo() {

}
