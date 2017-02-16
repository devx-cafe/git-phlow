package phlow

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

func TestSanitizeIssueToBranchName(t *testing.T) {
	testsToRun := [6]testCase{
		{issue: 12, branchName: "work on iss", expected: "12-work-on-iss", casedesc: "Test replaces whitespaces with dash '-'"},
		{issue: 45, branchName: "Case SENsitivity", expected: "45-case-sensitivity", casedesc: "Test converts charecters to lowercase"},
		{issue: 15, branchName: ".branch name", expected: "15-branch-name", casedesc: "Test removes . prefix"},
		{issue: 220, branchName: "^^..:~:name", expected: "220---name", casedesc: "removes ASCII control characters"},
		{issue: 2735, branchName: "name/", expected: "2735-name", casedesc: "test removes end / "},
		{issue: 234567, branchName: ".NAME.is\"dotted", expected: "234567-name-isdotted", casedesc: "test removes backslash"},
	}

	Convey("Test Sanitize Function", t, func() {

		for _, currentTest := range testsToRun {

			Convey(currentTest.casedesc, func() {

				actualName := SanitizeIssueToBranchName(currentTest.issue, currentTest.branchName)
				So(actualName, ShouldEqual, currentTest.expected)
			})
		}
	})
}
