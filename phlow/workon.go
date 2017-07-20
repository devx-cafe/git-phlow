package phlow

import (
	"fmt"
	"os"

	"strconv"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
)

//WorkOn ...
func WorkOn(issue int) {
	ui.PhlowSpinner.Start("Setting up workspace")

	defer ui.PhlowSpinner.Stop()
	if err := githandler.Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	branchInfo, err := githandler.Branch()
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	if plugins.IssueFromBranchName(branchInfo.Current) == issue {
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.Format.Branch(branchInfo.Current))
		return
	}

	for _, branch := range branchInfo.List {
		if plugins.IssueFromBranchName(branch) == issue {
			if err = githandler.CheckOut(branch); err != nil {
				fmt.Println(err)
			}
			ui.PhlowSpinner.Stop()
			fmt.Fprintf(os.Stdout, "Resuming to workspace:  %s \n", ui.Format.Branch(branch))
			return
		}
	}

	//Get list of gh issues
	gitHubIssues, err := plugins.GitHub.GetIssues()
	if err != nil {
		fmt.Println(err)
		return
	}
	//Get the default branch - eg. master
	defaultBranch, err := plugins.GitHub.Default()
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
			ui.PhlowSpinner.Stop()
			fmt.Fprintf(os.Stdout, "Created workspace:  %s \n", ui.Format.Branch(name))

			//Set labels and Assignee
			UpdateIssue(issue)
			return
		}
	}

	fmt.Println("No matching issues")
}

//UpdateIssue ...
//Set Label and assignee on a GitHub issue
func UpdateIssue(issue int) {
	//Retrieve token
	user := githandler.ConfigGet("user", "phlow")

	if _, err := plugins.GitHub.SetLabel(plugins.PhlowLabels["Status - in progress"].Title, issue); err != nil {
		fmt.Println(err)
	}

	if err := plugins.GitHub.SetAssignee(user, issue); err != nil {
		fmt.Println(err)
	}

	is := strconv.Itoa(issue)
	fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.Format.Issue(is))
	fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.Format.Label.G4Move(plugins.PhlowLabels["Status - in progress"].Title))
	fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.Format.Assignee(user))
	fmt.Println("----------------------------------")
	return
}
