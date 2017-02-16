package phlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/gitwrapper"
	"github.com/praqma/git-phlow/plugins"
	"gopkg.in/libgit2/git2go.v25"
)

//WorkOn ...
//Workon function
func WorkOn(issueFromUser int, plugin plugins.Plugin, git gitwrapper.Giter) {

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
	if plugin.IsSignedIn() {
		fmt.Fprintln(os.Stdout, "You have not generated a token, or signed in to your repo, run gen cmd ")
	}

	if fetchOutput, err := git.Fetch().FetchFromOrigin(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		fmt.Fprintln(os.Stdout, fetchOutput)
	}

	//Get local branches and map key:issue number, value: branchname
	branches, _ := git.Branch().ListBranches()
	mappedBranches := getBranchesAsMap(branches)

	//Add option to rework on already delivered branches

	if mappedBranches[issueFromUser] != "" {
		SwitchOrReworkExistingBranch(mappedBranches[issueFromUser], git)
	} else {
		CheckoutNewBranchFromPluginIssue(issueFromUser, plugin, git)
	}
}

func CheckoutNewBranchFromPluginIssue(issueFromUser int, plugin plugins.Plugin, git gitwrapper.Giter) {

	pluginIssues := plugin.ListIssues()     //GetIssues
	defaultBranch := plugin.DefaultBranch() //Get default branch

	if pluginIssues[issueFromUser] != "" {
		//Issue matches with issueNumberInput
		newBranchName := SanitizeIssueToBranchName(issueFromUser, pluginIssues[issueFromUser])
		output, err := git.Checkout().CheckoutNewBranchFromOrigin(newBranchName, defaultBranch)
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		} else {
			fmt.Fprintln(os.Stdout, output)
		}
	} else {
		fmt.Fprintln(os.Stdout, "Issue does not exist in your repository")
	}
}

func SwitchOrReworkExistingBranch(branchName string, git gitwrapper.Giter) {
	//Branch is already created - do checkout
	_, err := git.Checkout().Checkout(branchName)

	if err == nil {
		//No file conflicts at checkout
		fmt.Fprintln(os.Stdout, "branch "+branchName+" already created from issue ")
		fmt.Fprintln(os.Stdout, "Switching to branch branchMap[issuenumber]")
	} else {
		fmt.Fprint(os.Stdout, err)
	}
}


//Helper methods
func getBranchesAsMap(branches []string) map[int]string {
	mappings := make(map[int]string)
	for _, branch := range branches {
		var tmp = branch
		num, err := extractIssueNumber(branch)
		if err == nil {
			mappings[num] = tmp
		}
	}
	return mappings
}

func extractIssueNumber(str string) (int, error) {
	issueString := strings.Split(str, "-")[0]
	return strconv.Atoi(issueString)
}
