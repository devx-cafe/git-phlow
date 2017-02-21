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
func WorkOn(issue int) {
	//Before check if i am logged in

	if err := githandler.Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	branchInfo, err := githandler.Branch("list")
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	for _, branch := range branchInfo.List {
		if strings.HasPrefix(branch, strconv.Itoa(issue)) {
			githandler.CheckOut(branch, false)
			return
		}
	}

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

			fmt.Fprintf(os.Stdout, "branch '%s' created and checkout out for work", name)
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
}
