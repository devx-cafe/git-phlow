package phlow

import (
	"fmt"
	"os"
	"strconv"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
)

//IssueList ...
//List open issues from GitHub
func IssueList() {

	ui.PhlowSpinner.Start("")
	issues, err := plugins.GitHub.Issue.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	ui.PhlowSpinner.Stop()

	//Nested function for finding user issues
	var userIssue = func(issue []plugins.AssigneeIssue) bool {
		user := githandler.ConfigGet("user", "phlow")
		for _, u := range issue {
			if u.Login == user {
				return true
			}
		}
		return false
	}
	//Nested print function - only used for pretty output
	var printIssue = func(issue plugins.Issues) {
		issStr := strconv.Itoa(issue.Number)

		fmt.Print(ui.Format(issStr + ": ").Bold)
		fmt.Print(issue.Title)

		for _, label := range issue.Labels {

			fmt.Print(" " + ui.Format(label.Name).Label.FByG(plugins.GroupID(label.Name)))
		}

		for _, user := range issue.Assignees {
			fmt.Print(" " + ui.Format(user.Login).Assignee)
		}
		fmt.Print(" " + ui.Format(issue.Milestone.Title).MileStone)

		fmt.Println()
	}

	fmt.Println(ui.Format("# Issues").MileStone)

	for _, issue := range issues {
		assignees := issue.Assignees
		//If mine is true we print on issues assigned to a user
		if options.GlobalFlagMine {
			if userIssue(assignees) {
				printIssue(issue)
			}
		} else {
			printIssue(issue)
		}
	}

}
