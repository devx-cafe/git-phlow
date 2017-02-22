package phlow

import (
	"fmt"
	"strings"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"os"
)

//LocalDeliver ...
func LocalDeliver() {

	branchInfo, _ := githandler.Branch()
	dfBranch, _ := plugins.GetDefaultBranch(plugins.RepoUrl)

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == dfBranch) {
		fmt.Printf("You cannot deliver: %s", branchInfo.Current)
		return
	}
	fmt.Fprintf(os.Stdout, "Checking out default branch '%s' \n", dfBranch)
	//Checkout default branch: master
	if err := githandler.CheckOut(dfBranch, false); err != nil {
		fmt.Println(err)
		return
	}
	//Pull rebase latest changes
	fmt.Fprintln(os.Stdout, "Trying to pull latest changes")
	output, err := githandler.Pull()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	fmt.Fprintf(os.Stdout, "Merging changes from branch '%s' into branch '%s' \n", branchInfo.Current, dfBranch)
	//Merge feature branch into default
	if err := githandler.Merge(branchInfo.Current); err != nil {
		fmt.Println(err)
	}
	//Rename default branch to delivered
	githandler.BranchRename(branchInfo.Current)

	//Push changes to github
	fmt.Fprintf(os.Stdout, "Pushing changes to remote '%s' \n", dfBranch)
	output, err = githandler.Push("", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	fmt.Printf("Branch '%s' fearlessly delivered to '%s'\n", branchInfo.Current, dfBranch)
	
}
