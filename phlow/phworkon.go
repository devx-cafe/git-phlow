package phlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/gitwrapper"
	"github.com/praqma/git-phlow/plugins"
	"errors"
)

//WorkOn ...
//Creates a branch from origin/defaultbranch named after an issue
// from the plugin and does a Checkout on that branch
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
	if plugins.IsSignedIn() {
		fmt.Fprintln(os.Stdout, "You have not generated a token, or signed in to your repo, run gen cmd ")
	}

	if fetchOutput, err := git.Fetch().FetchFromOrigin(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		fmt.Fprintln(os.Stdout, fetchOutput)
	}

	//Get local branches and map key:issue number, value: branchName
	branches, _ := git.Branch().ListBranches()
	mappedBranches := getBranchesAsMap(branches)

	//Add option to rework on already delivered branches

	if mappedBranches[issueFromUser] != "" {
		SwitchOrReworkExistingBranch(mappedBranches[issueFromUser], git)
	} else {
		CheckoutNewBranchFromPluginIssue(issueFromUser, plugin, git)
	}
}

//CheckoutNewBranchFromPluginIssue ...
//Checkout a new branch from a plugin issue if the issue exists
func CheckoutNewBranchFromPluginIssue(issueFromUser int, plugin plugins.Plugin, git gitwrapper.Giter) error {

	pluginIssues := plugin.ListIssues()               //GetIssues
	defaultBranchFromRemote := plugin.DefaultBranch() //Get default branch

	if pluginIssues[issueFromUser] != "" {
		//Issue matches with issueNumberInput
		newBranchName := SanitizeIssueToBranchName(issueFromUser, pluginIssues[issueFromUser])
		output, err := git.Checkout().CheckoutNewBranchFromOrigin(newBranchName, defaultBranchFromRemote)
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			return err
		}
		fmt.Fprintln(os.Stdout, output)

	} else {
		err := errors.New("Issue does not exist in your repository")
		fmt.Fprintln(os.Stdout, err)
		return err
	}
	return nil
}

//SwitchOrReworkExistingBranch ...
//Switches to an already existing branch if working on same issue
func SwitchOrReworkExistingBranch(branchName string, git gitwrapper.Giter) error {
	//Branch is already created - do checkout
	_, err := git.Checkout().Checkout(branchName)

	if err == nil {
		//No file conflicts at checkout
		fmt.Fprintf(os.Stdout, "branch: '%s' already created \n", branchName)
		fmt.Fprintf(os.Stdout, "Switching to branch: '%s' and resuming phlow work \n", branchName)
	} else {
		fmt.Fprint(os.Stdout, err)
		return err
	}
	return nil
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
