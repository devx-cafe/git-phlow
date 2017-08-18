package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/executor"
)

//MakeAliasCaller ...
func MakeAliasCaller() {
	MakeAlias()
}

//MakeAlias ...
func MakeAlias() {
	git := githandler.Git{Run: executor.RunGit}
	aliases := make(map[string]string)
	aliases["alias.wrapup"] = "phlow wrapup"
	aliases["alias.workon"] = "phlow workon"
	aliases["alias.deliver"] = "phlow deliver"
	aliases["alias.cleanup"] = "phlow cleanup"
	aliases["alias.web"] = "phlow web"
	aliases["alias.issues"] = "phlow issues"

	for group, value := range aliases {

		str, _ := git.Config("--global", "--get", group)
		if str == "" {
			fmt.Printf("Creating alias %s \n", ui.Format.Alias(group))
			git.Config("--global", group, value)
		} else {
			fmt.Printf("Alias %s already exists \n", ui.Format.Alias(group))
		}
	}
}
