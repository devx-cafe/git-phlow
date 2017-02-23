package phlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
)

//WorkOn ...
func WorkOn(issue int) {

	fmt.Fprintln(os.Stdout, "Fetching changes from remote")
	if err := githandler.Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	branchInfo, err := githandler.Branch()
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	fmt.Fprintln(os.Stdout, "Locating existing issue branches")
	if GetIssueFromBranch(branchInfo.Current) == issue {
		fmt.Fprintf(os.Stdout, "You are already on branch '%s'\n", branchInfo.Current)
		return
	}

	for _, branch := range branchInfo.List {
		if GetIssueFromBranch(branch) == issue {
			if err = githandler.CheckOut(branch); err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Switched to branch '%s' \n", branch)
			return
		}
	}

	fmt.Fprintln(os.Stdout, "No 'local' issue branches found. Searching on github")
	info, err := plugins.GetOpenIssues(plugins.RepoURL)
	if err != nil {
		fmt.Println(err)
	}

	//Getting default branch
	dfBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)

	for _, iss := range info {
		if iss.Number == issue {
			name := plugins.BranchNameFromIssue(issue, iss.Title)
			if err = githandler.CheckoutNewBranchFromRemote(name, dfBranch); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(os.Stdout, "branch '%s' created and checked out \n", name)

			//Retrieve token
			token := githandler.ConfigGet("token", "phlow")
			user := githandler.ConfigGet("user", "phlow")

			if _, labelErr := plugins.SetLabel(plugins.LabelStatusInProgress, plugins.RepoURL, token, issue); labelErr != nil {
				fmt.Println(labelErr)
			}

			if assigneeArr := plugins.SetAssignee(user, plugins.RepoURL, token, issue); err != nil {
				fmt.Println(assigneeArr)
			}
			fmt.Fprintf(os.Stdout, "Issue updated with label '%s' and assignee '%s' \n", plugins.LabelStatusInProgress, user)
			return
		}
	}

	fmt.Println("No 'remote' issues matches you input")
}

//GetIssueFromBranch ...
func GetIssueFromBranch(branch string) int {
	iss, err := strconv.Atoi(strings.Split(branch, "-")[0])
	if err != nil {
		return -1
	}
	return iss
}
