package phlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/gitwrapper"
	"github.com/praqma/git-phlow/plugins"
)

//WorkOn ...
//Workon function
func WorkOn(issueNumber int, plugin plugins.Plugin, git gitwrapper.Giter) {

	//Verify is located in initialized repository
	if err := git.Status().Status() != nil; err {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	//Verifying if repository has a remote
	if !git.Fetch().HasRemote() {
		fmt.Fprintln(os.Stdout, "Fatal error: No origin - Repository need to have an origin")
		return
	}

	//Verify token is generated and we are signed into plugin
	//if plugin.IsSignedIn(){
	//	fmt.Fprintln(os.Stdout, "You have not generated a token, or signed in to your repo, run gen cmd ")
	//}

	if fetchOutput, err := git.Fetch().FetchFromOrigin(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		fmt.Fprintln(os.Stdout, fetchOutput)
	}

	//Get local branches and map key:issue number, value: branchname
	branches, _ := git.Branch().ListBranches()
	mappedBranches := getBranchIssues(branches)

	//Add option to rework on already delivered branches

	if mappedBranches[issueNumber] != "" {
		//Branch is already created - do checkout
		_, err := git.Checkout().Checkout(mappedBranches[issueNumber])

		if err == nil {
			//No file conflicts at checkout
			fmt.Fprintln(os.Stdout, "branch "+mappedBranches[issueNumber]+" already created from issue ")
			fmt.Fprintln(os.Stdout, "Switching to branch branchMap[issuenumber]")
		} else {
			fmt.Fprint(os.Stdout, err)
		}
	} else {
		pluginIssues := plugin.ListIssues()

		if pluginIssues[issueNumber] != "" {
			//newBranchName := SanitizeIssueToBranchName(issueNumber, pluginIssues[issueNumber])
			//branch, err := git.Branch().CreateBranch(newBranchName)

		}
	}
}

func updateOriginAndContinue(git gitwrapper.Giter) {

	if fetch := git.Fetch(); fetch.HasRemote() {
		fetch.FetchFromOrigin()
		fmt.Fprintln(os.Stdout, "Fetching remote branches")
	} else {
		fmt.Fprintln(os.Stderr, "")

	}
}

func getBranchIssues(branches []string) map[int]string {
	mappings := make(map[int]string)
	for _, branch := range branches {
		var tmp = branch
		num, err := extractIssue(branch)
		if err == nil {
			mappings[num] = tmp
		}
	}
	return mappings
}

func extractIssue(str string) (int, error) {
	issueString := strings.Split(str, "-")[0]
	return strconv.Atoi(issueString)
}
