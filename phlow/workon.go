package phlow

import (
	"fmt"
	"os"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
)

//WorkOn ...
func WorkOn(issue int) {

	ui.PhlowSpinner.Start("Fetching form remote")
	if err := githandler.Fetch(); err != nil {
		fmt.Println(err)
		return
	}
	ui.PhlowSpinner.Stop()
	fmt.Println("Updates fetched")

	branchInfo, err := githandler.Branch()
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	fmt.Fprintln(os.Stdout, "Locating existing issue branches")
	if plugins.IssueFromBranchName(branchInfo.Current) == issue {
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.BranchFormat(branchInfo.Current))
		return
	}

	for _, branch := range branchInfo.List {
		if plugins.IssueFromBranchName(branch) == issue {
			if err = githandler.CheckOut(branch); err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Switched to branch %s \n", ui.BranchFormat(branch))
			return
		}
	}
	fmt.Fprintf(os.Stdout, "No local %s found. Searching github \n", ui.Bold("issue-branches"))

	//Get list of github issues
	gitHubIssues, err := plugins.GitHub.Issue.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	//Get the default branch - eg. master
	defaultBranch, err := plugins.GitHub.Branch.Default()
	if err != nil {
		fmt.Println(err)
	}

	//Loop through all issues verifying the work-on issue exists
	for _, iss := range gitHubIssues {
		if iss.Number == issue {
			name := plugins.BranchNameFromIssue(issue, iss.Title)
			if err = githandler.CheckoutNewBranchFromRemote(name, defaultBranch); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(os.Stdout, "branch %s created and checked out \n", ui.BranchFormat(name))

			//Set labels and Assignee
			UpdateIssue(issue)
			return
		}
	}

	fmt.Println("No issues matches you input")
}

//UpdateIssue ...
//Set Label and assignee on a GitHub issue
func UpdateIssue(issue int) {
	//Retrieve token
	user := githandler.ConfigGet("user", "phlow")

	if _, err := plugins.GitHub.Label.Set(plugins.PhlowLabels["Status - in progress"].Title, issue); err != nil {
		fmt.Println(err)
	}

	if err := plugins.GitHub.Assignee.Set(user, issue); err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.IssueFormat(issue))
	fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.LabelFormat(plugins.PhlowLabels["Status - in progress"].Title))
	fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.AssigneeFormat(user))
	fmt.Println("----------------------------------")
	return
}
