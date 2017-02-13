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
			fmt.Fprintln(os.Stdout, "branch "+branchMappings[issueNumber]+" already created from issue ", )
			fmt.Fprintln(os.Stdout, "Switching to branch branchMap[issuenumber]")
		} else {
			fmt.Fprint(os.Stdout, err)
		}

	} else {
		//Creating new issue-branch
		branch, err := git.Branch().CreateBranch(strconv.Itoa(issueNumber) + "-issue-default")

		if err == nil {
			//issue branch successfully created
			fmt.Fprintln(os.Stdout, "Branch "+branch+" successfully created")
			_, err := git.Checkout().Checkout(branch)

			if err == nil {
				fmt.Fprintln(os.Stdout, "Switched to branch: "+branch)
			} else {
				fmt.Fprintln(os.Stdout, "Could not checkout branch")
			}
		} else {
			fmt.Fprintln(os.Stdout, "Could not create branch "+branch)

		}
	}
}
func updateOriginAndContinue(git gitwrapper.Giter) {

	if fetch := git.Fetch(); fetch.HasRemote() {
		fetch.FetchFromOrigin()
		fmt.Fprintln(os.Stdout, "Fetching remote branches")
	} else {
		fmt.Fprintln(os.Stderr, "Working on local copy")
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
