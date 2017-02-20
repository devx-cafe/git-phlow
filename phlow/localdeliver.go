package phlow

import (
	"fmt"
	"strings"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
)

//FearlessDeliver ...
func FearlessDeliver() {

	//Prechecks - status

	branchInfo, _ := githandler.Branch("current")
	dfBranch, _ := plugins.GetDefaultBranch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == dfBranch) {
		fmt.Printf("You cannot deliver: %s", branchInfo.Current)
		return
	}

	//Checkout default branch: master
	if err := githandler.CheckOut(dfBranch, false); err != nil {
		fmt.Println(err)
		return
	}
	//Pull rebase latest changes
	output, err := githandler.Pull()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	//Merge feature branch into default
	if err := githandler.Merge(branchInfo.Current); err != nil {
		fmt.Println(err)
	}
	//Rename default branch to delivered
	githandler.BranchRename(branchInfo.Current)

	//Push changes to github
	output, err = githandler.Push("", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	fmt.Printf("Branch '%s' fearlessly delivered to '%s' - A small celebration is in order \n", branchInfo.current, dfBranch)
}
