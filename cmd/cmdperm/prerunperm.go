package cmdperm

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/platform"
)

//RequiredAuthentication ...
//Validates if the user has logged in before running the command
func RequiredAuthentication() {
	conf := platform.DefaultConfiguration()
	token := conf.Get(platform.PhlowToken)
	user := conf.Get(platform.PhlowUser)

	if token == "" || user == "" {
		fmt.Printf("Please run %s to connect to github \n", ui.Format.Bold("auth"))
		os.Exit(0)
	}
}

//RequiredCurDirRepository ...
//Validates if the command is runnign in a git repository
func RequiredCurDirRepository() {
	if err := githandler.Status(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(0)
	}
}
