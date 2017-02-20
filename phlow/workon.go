package phlow

import "fmt"

//WorkOn ...
func WorkOn(newBranch string) {
	//Before check if i am logged in

	//Move as before check
	if err := Status(); err != nil {
		fmt.Println("fatal: not in git repository")
		return
	}

	if err := Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	branchinfo, err := Branch("list")
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}
	for _, branch := range branchinfo.list {
		if branch == newBranch {
			CheckOut(branch, false)
			return
		}
	}

	if err := CheckOut(newBranch, true); err != nil {
		fmt.Println(err)
		return
	}

	//Set assignee
	if err := SetAssignee(); err != nil {
		fmt.Println(err)
	}
	//SetLabel
	if err := SetLabel("Status - in progress"); err != nil {
		fmt.Println(err)
	}
}
