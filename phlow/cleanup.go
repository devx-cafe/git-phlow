package phlow

import (
	"fmt"

	"github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/githandler"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/setting"
	"github.com/code-cafe/git-phlow/ui"
)

//CleanCaller ...
func CleanCaller(ini string) {
	conf := setting.NewProjectStg(ini)
	Clean(conf)
}

//Clean ...
//deletes all the delivered branches
func Clean(conf *setting.ProjectSetting) {
	git := githandler.Git{Run: executor.RunGit}

	out, err := git.Branch("-a")
	if err != nil {
		fmt.Println(err)
	}

	local, remote := githandler.Delivered(githandler.AsList(out), conf.Remote)

	for _, branch := range local {
		deleteFlag := "-d"

		if options.GlobalFlagForce {
			deleteFlag = "-D"
		}

		_, err = git.Branch(deleteFlag, branch)
		if err != nil {
			fmt.Printf("Could not delete branch %s \n", branch)
		} else {
			fmt.Printf("Deleted local branch %s \n", ui.Format.Branch(branch))
		}
	}

	//Remote branches should be deleted as well
	if !options.GlobalFlagLocal {
		for _, branch := range remote {
			_, err = git.Push(branch, conf.Remote, "--delete", branch)
			if err != nil {
				fmt.Printf("Could not delete branch %s \n", branch)
			} else {
				fmt.Printf("Deleted remote branch %s \n", ui.Format.Branch(branch))
			}
		}
		git.Fetch("--prune")
	}

	if options.GlobalFlagLocal {
		if len(remote) != 0 {
			fmt.Println("Remote branches ready to be deleted: ")
			fmt.Println(remote)
		}
	}
}
