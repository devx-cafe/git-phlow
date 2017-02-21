package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"strings"
	"strconv"
	"os"
)

//WorkOn ...
func WorkOn(issue int, verbose bool) {

	printVerbose("Fetching changes from remote", verbose)
	if err := githandler.Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	branchInfo, err := githandler.Branch("list")
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	printVerbose("Locating existing issue branches", verbose)
	for _, branch := range branchInfo.List {
		if GetIssueFromBranch(branch) == issue {
			if err := githandler.CheckOut(branch, false); err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Switched to branch '%s' \n", branch)
			return
		}
	}

	printVerbose("No 'local' issue branches found. Searching on github", verbose)
	info, err := plugins.GetOpenIssues(plugins.RepoUrl)
	if err != nil {
		fmt.Println(err)
	}

	for _, iss := range info {
		if iss.Number == issue {
			name := plugins.BranchNameFromIssue(issue, iss.Title)
			if err := githandler.CheckOut(name, true); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(os.Stdout, "branch '%s' created and checked out", name)
			return
		}
	}

	//Set assignee
	//if err := plugins.SetAssignee(); err != nil {
	//	fmt.Println(err)
	//}
	////SetLabel
	//if err := plugins.SetLabel("Status - in progress"); err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println("No 'remote' issues matches you input")
}

func printVerbose(message string, verbose bool) {
	if verbose {
		fmt.Fprintln(os.Stdout, message)
	}
}

func GetIssueFromBranch(branch string) int {
	iss, err := strconv.Atoi(strings.Split(branch, "-")[0])
	if err != nil {
		return -1
	}
	return iss
}
