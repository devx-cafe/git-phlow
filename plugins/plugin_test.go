package plugins

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"strconv"
)

type testCase struct {
	issue      int
	branchName string
	expected   string
	casedesc   string
}

/* Branch naming conventions
 * - cant begin with '.'
 * - cant have double dot '..'
 * - cant have chars '^', '~', ':'
 * - end with backslash /
 * - end with ".lock"
 * - contain \
 */

func TestBranchNameFromIssue(t *testing.T) {
	testsToRun := [7]testCase{
		{issue: 12, branchName: "work on iss", expected: "12-work-on-iss", casedesc: "Test replaces whitespaces with dash '-'"},
		{issue: 45, branchName: "Case SENsitivity", expected: "45-case-sensitivity", casedesc: "Test converts charecters to lowercase"},
		{issue: 15, branchName: ".branch name", expected: "15-branch-name", casedesc: "Test removes . prefix"},
		{issue: 220, branchName: "^^..:~:name", expected: "220-name", casedesc: "removes ASCII control characters"},
		{issue: 2735, branchName: "name/", expected: "2735-name", casedesc: "test removes end / "},
		{issue: 234567, branchName: ".NAME.is\"dotted", expected: "234567-name-is-dotted", casedesc: "test removes backslash"},
		{issue: 672, branchName: "add big data blog /", expected: "672-add-big-data-blog", casedesc: "remove forward slash"},
	}

	Convey("Running tests on 'BranchNameFromIssue'", t, func() {

		for _, currentTest := range testsToRun {

			Convey(currentTest.casedesc, func() {
				actualName := BranchNameFromIssue(strconv.Itoa(currentTest.issue), currentTest.branchName)
				t.Log(currentTest.branchName)
				So(actualName, ShouldEqual, currentTest.expected)
			})
		}
	})
}

func TestStringConcat(t *testing.T) {

	Convey("Running tests on 'effecientStringConcat' function ", t, func() {

		Convey("Test variable inputs get concatenated correctly", func() {
			var expectedLong = "created 'something' new"
			var actualLong = efficientConcatString("created '", "something", "' new")
			So(actualLong, ShouldEqual, expectedLong)
		})

		Convey("Test funny signs gets concatenated as well", func() {
			var expectedShort = "j$¢‰¿≈¯¯¯"
			var actualShort = efficientConcatString("j$¢‰¿≈", "¯¯¯")
			So(expectedShort, ShouldEqual, actualShort)
		})
	})
}

func TestGetIssueFromBranch(t *testing.T) {
	Convey("Running tests on 'GetIssuesFromBranch'", t, func() {

		Convey("GetIssueSFromBranch should return 1", func() {

			i := IssueFromBranchName("39-enable---Add-sign-in-function")
			So(i, ShouldEqual, "39")
		})
	})
}
