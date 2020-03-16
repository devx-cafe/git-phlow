package phlow

import (
	"fmt"

	"github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/githandler"
	"github.com/code-cafe/git-phlow/setting"
)

//Bootstrap ...
//Creates a new .gitconfig file with a default configuration
func Bootstrap() {

	git := githandler.Git{Run: executor.RunGit}

	local, err := git.RevParse("--show-toplevel")
	if err != nil {
		panic(err)
	}

	setting.BootstrapPhlowConfig(local, "master")
}

//Show ...
//Shows the listed configuration
func Show(args []string) {

	if len(args) >= 1 {
		conf := setting.NewProjectStg(args[0])
		fmt.Println(conf.ToString())
		return
	}
	conf := setting.NewProjectStg("")
	fmt.Println(conf.ToString())
}
