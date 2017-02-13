package phlowimpl

import (
	"github.com/praqma/git-phlow/gitwrapper"
	"strings"
	"strconv"
	"os"
	"fmt"
)

//WorkOn
func WorkOn(issueNumber int) {

	git := gitwrapper.InitGit()
	branchMappings := make(map[int]string)
	branches, err := git.Branch().ListBranches()

	if err != nil {

		//Check if repo
		//check if git is in path

		fmt.Fprintln(os.Stdout, "")
		return
	}

	getIssues(branches, branchMappings)

	if branchMappings[issueNumber] != "" {
		//Branch is already created - do checkout
		git.Checkout().Checkout(branchMappings[issueNumber])
		fmt.Fprintln(os.Stdout, "branch %s already created from issue %s", branchMappings[issueNumber])
		fmt.Fprintln(os.Stdout, "Switching to branch branchMap[issuenumber]")
	}

	//issues getIssues

	//br, err := git.Branch().CreateBranch("name")

	if err != nil {
		//cannot create branch

	}

}

func getIssues(branches []string, mappings map[int]string) {

	for _, branch := range branches {
		var tmp = branch
		num, err := extractIssue(branch)
		if err == nil {
			mappings[num] = tmp
		}
	}
}

func extractIssue(str string) (int, error) {
	issueString := strings.Split(str, "-")[0]
	return strconv.Atoi(issueString)
}
