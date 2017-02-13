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

	updateOriginAndContinue(git)

	branches, err := git.Branch().ListBranches()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	getIssues(branches, branchMappings)

	if branchMappings[issueNumber] != "" {
		//Branch is already created - do checkout
		_, err := git.Checkout().Checkout(branchMappings[issueNumber])

		if err == nil {
			//No file conflicts at checkout
			fmt.Fprintln(os.Stdout, "branch %s already created from issue %s", branchMappings[issueNumber])
			fmt.Fprintln(os.Stdout, "Switching to branch branchMap[issuenumber]")
		} else {
			fmt.Fprint(os.Stdout, err)
		}

	} else {
		//Creating new issue-branch
		str, err := git.Branch().CreateBranch(strconv.Itoa(issueNumber) + "issue-default")

		if err == nil {
			//issue branch successfully created
			fmt.Fprint(os.Stdout, "Branch %s successfully created", str)
			_, err := git.Checkout().Checkout(str)

			if err == nil {
				fmt.Fprint(os.Stdout, "Switched to branch "+str)
			} else {
				fmt.Fprintln(os.Stdout, "Could not checkout branch")
			}
		} else {
			fmt.Fprintln(os.Stdout, "Could not create branch "+str)

		}
	}
}
func updateOriginAndContinue(git gitwrapper.Giter) {

	if fetch := git.Fetch(); fetch.HasRemote() {
		fetch.FetchFromOrigin()
		fmt.Fprintln(os.Stdout, "Fetching remote branches")
	} else {
		fmt.Fprint(os.Stdout, "No remote found, working on local copy")
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
