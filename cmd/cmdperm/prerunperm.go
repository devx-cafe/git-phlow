package cmdperm

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/executor"
)

//RequiredCurDirRepository ...
//Validates if the command is runnign in a git repository
func RequiredCurDirRepository() {
	git := githandler.Git{Run: executor.RunGit}
	if _, err := git.Status(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(0)
	}
}
