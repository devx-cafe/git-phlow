package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
)

//WorkOn ...
func WorkOn(newBranch string) {
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
		if branch == newBranch {
			githandler.CheckOut(branch, false)
			return
		}
	}

	if err := githandler.CheckOut(newBranch, true); err != nil {
		fmt.Println(err)
		return
	}

	//Set assignee
	if err := plugins.SetAssignee(); err != nil {
		fmt.Println(err)
	}
	//SetLabel
	if err := plugins.SetLabel("Status - in progress"); err != nil {
		fmt.Println(err)
	}
}
