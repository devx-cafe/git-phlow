package phlow

import (
	"fmt"
	"strings"
)

//Deliver ...
func Deliver() {
	//prechecks - status

	branchInfo, _ := Branch("current")
	dfBranch, _ := GetDefaultBranch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.current, "delivered/") || (branchInfo.current == dfBranch) {
		fmt.Printf("You cannot deliver: %s", branchInfo.current)
		return
	}

	output, err := Push(branchInfo.current, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	if err := BranchRename(branchInfo.current); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Branch '%s' is now delivered", branchInfo.current)
}
