package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
)

//Deliver ...
func Deliver() {

	branchInfo, _ := githandler.Branch()
	dfBranch, _ := plugins.GetDefaultBranch(plugins.RepoUrl)

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == dfBranch) {
		fmt.Printf("You cannot deliver: %s", branchInfo.Current)
		return
	}

	output, err := githandler.Push(branchInfo.Current, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	if err := githandler.BranchRename(branchInfo.Current); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Branch '%s' is now delivered", branchInfo.Current)
}
