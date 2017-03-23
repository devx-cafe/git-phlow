package phlow

import (
	"fmt"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/ui"
)

//Clean ...
func Clean(remoteName string) {

	local, remote := githandler.BranchDelivered(remoteName)
	var err error

	for _, branch := range local {
		if _, err = githandler.BranchDelete(branch, "", false, options.GlobalFlagForce); err != nil {
			fmt.Printf("Could not delete branch %s \n", branch)
		} else {
			fmt.Printf("Local branch %s deleted \n", ui.Format.Branch(branch))
		}
	}

	if !options.GlobalFlagLocal {
		for _, branch := range remote {
			if _, err = githandler.BranchDelete(branch, remoteName, true, false); err != nil {
				fmt.Printf("Could not delete branch %s \n", branch)
			} else {
				fmt.Printf("Remote branch %s deleted \n", ui.Format.Branch(branch))
			}
		}
		githandler.FetchPrune()
	}

	if options.GlobalFlagLocal {
		if len(remote) != 0 {
			fmt.Println("Remote branches ready to be deleted: ")
			fmt.Println(remote)
		}
	}
}
