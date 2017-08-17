package phlow

import (
	"fmt"
	"strings"

	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/executor"
)

//WrapUp ...
func WrapUp() {
	git := githandler.Git{Run: executor.RunGit}

	//Add all files to index
	_, err := git.Add("--all")
	if err != nil {
		fmt.Println("Project files could not be added: " + err.Error())
		return
	}

	out, err := git.Branch("-a")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Retrieve branch info - current branch
	info := githandler.AsList(out)

	var commitMessage string

	if options.GlobalFlagForceMessage != "" {
		commitMessage = "close #" + strings.Split(info.Current, "-")[0] + " " + options.GlobalFlagForceMessage
	} else {
		commitMessage = "close #" + strings.Replace(info.Current, "-", " ", -1)
	}

	_, err = git.Commit("-m", commitMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(os.Stdout, commitMessage)
}
