package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/options"
)

//Deliver ...
func Deliver() {

	branchInfo, _ := githandler.Branch()
	dfBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)

	githandler.Fetch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == dfBranch) {
		fmt.Printf("Could not deliver: %s", branchInfo.Current)
		return
	}

	output, err := githandler.PushRename(branchInfo.Current, dfBranch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	if err := githandler.BranchRename(branchInfo.Current); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Branch %s  is now delivered \n", options.BranchFormat(branchInfo.Current))
}
