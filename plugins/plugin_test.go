package plugins_test

import (
	. "github.com/praqma/git-phlow/plugins"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("Plugin", func() {

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

	Describe("Branch name from issue", func() {
		testsToRun := [7]testCase{
			{issue: 12, branchName: "work on iss", expected: "12-work-on-iss", casedesc: "Test replaces whitespaces with dash '-'"},
			{issue: 45, branchName: "Case SENsitivity", expected: "45-case-sensitivity", casedesc: "Test converts charecters to lowercase"},
			{issue: 15, branchName: ".branch name", expected: "15-branch-name", casedesc: "Test removes . prefix"},
			{issue: 220, branchName: "^^..:~:name", expected: "220-name", casedesc: "removes ASCII control characters"},
			{issue: 2735, branchName: "name/", expected: "2735-name", casedesc: "test removes end / "},
			{issue: 234567, branchName: ".NAME.is\"dotted", expected: "234567-name-is-dotted", casedesc: "test removes backslash"},
			{issue: 672, branchName: "add big data blog /", expected: "672-add-big-data-blog", casedesc: "remove forward slash"},
		}

		Context("Names should follow format rules", func() {

			for _, currentTest := range testsToRun {

				It(currentTest.casedesc, func() {
					actualName := BranchNameFromIssue(strconv.Itoa(currentTest.issue), currentTest.branchName)

					Ω(actualName).Should(Equal(currentTest.expected))
				})
			}
		})
	})

})
