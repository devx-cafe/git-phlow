package phlowimpl

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testCase struct {
	input    string
	expected string
	casedesc string
}

/* Branch naming conventions
 * - cant begin with '.'
 * - cant have double dot '..'
 * - cant have chars '^', '~', ':'
 * - end with backslash /
 * - end with ".lock"
 * - contain \
 */
//ConvertToBranchName
func TestInputs(t *testing.T) {

	testsToRun := [6]testCase{
		{input: "12 work on iss", expected: "12-work-on-iss", casedesc: "Test replaces whitespaces with dash '-'"},
		{input: "45 Case SENsitivity", expected: "45-case-sensitivity", casedesc: "Test converts charecters to lowercase"},
		{input: ".branch name", expected: "branch-name", casedesc: "Test removes . prefix"},
		{input: "^^..:~:name", expected: "--name", casedesc: "removes ASCII control characters"},
		{input: "name/", expected: "name", casedesc: "test removes end / "},
		{input: ".NAME.is\"dotted", expected: "name-isdotted", casedesc: "test removes backslash"},
	}

	Convey("Test branch name sanitation", t, func() {

		for _, currentTest := range testsToRun {

			Convey(currentTest.casedesc, func() {

				actualName := ConvertToBranchName(currentTest.input)
				t.Log(actualName)

				So(actualName, ShouldEqual, currentTest.expected)
			})
		}
	})
}
