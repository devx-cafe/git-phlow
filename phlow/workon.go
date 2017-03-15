package phlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
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
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.BranchFormat(branchInfo.Current))
		return
	}

	for _, branch := range branchInfo.List {
		if GetIssueFromBranch(branch) == issue {
			if err = githandler.CheckOut(branch); err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Switched to branch %s \n", ui.BranchFormat(branch))
			return
		}
	}

	fmt.Fprintf(os.Stdout, "No local %s found. Searching github \n", ui.Bold("issue-branches"))
	info, err := plugins.GetOpenIssues(plugins.RepoURL)
	if err != nil {
		fmt.Println(err)
	}

	//Getting default branch
	dfBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)
	remoteInfo, _ := githandler.Remote(dfBranch)

	for _, iss := range info {
		if iss.Number == issue {
			name := plugins.BranchNameFromIssue(issue, iss.Title)
			if err = githandler.CheckoutNewBranchFromRemote(name, dfBranch); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(os.Stdout, "branch %s created and checked out \n", ui.BranchFormat(name))

			//Retrieve token
			token := githandler.ConfigGet("token", "phlow")
			user := githandler.ConfigGet("user", "phlow")

			if _, labelErr := plugins.SetLabel(plugins.LabelStatusInProgress, plugins.RepoURL, token, issue, remoteInfo); labelErr != nil {
				fmt.Println(labelErr)
			}

			if assigneeArr := plugins.SetAssignee(user, plugins.RepoURL, token, issue, remoteInfo); err != nil {
				fmt.Println(assigneeArr)
			}
			fmt.Fprintf(os.Stdout, "\nIssue %s updated:  \n", ui.IssueFormat(issue))
			fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.LabelFormat(plugins.LabelStatusInProgress))
			fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.AssigneeFormat(user))
			return
		}
	}

	fmt.Println("No issues matches you input")
}

//GetIssueFromBranch ...
func GetIssueFromBranch(branch string) int {
	iss, err := strconv.Atoi(strings.Split(branch, "-")[0])
	if err != nil {
		return -1
	}
	return iss
}
