package phlow

import (
	"fmt"
	"strings"
)

//FearlessDeliver ...
func FearlessDeliver() {

	//Prechecks - status

	branchInfo, _ := Branch("current")
	dfBranch, _ := GetDefaultBranch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.current, "delivered/") || (branchInfo.current == dfBranch) {
		fmt.Printf("You cannot deliver: %s", branchInfo.current)
		return
	}

	//Checkout default branch: master
	if err := CheckOut(dfBranch, false); err != nil {
		fmt.Println(err)
		return
	}
	//Pull rebase latest changes
	output, err := Pull()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	//Merge feature branch into default
	if err := Merge(branchInfo.current); err != nil {
		fmt.Println(err)
	}
	//Rename default branch to delivered
	BranchRename(branchInfo.current)

	//Push changes to github
	output, err = Push("", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	fmt.Printf("Branch '%s' fearlessly delivered to '%s' - A small celebration is in order \n", branchInfo.current, dfBranch)
}
