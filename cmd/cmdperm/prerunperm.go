package cmdperm

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/setting"
)

//RequiredAuthentication ...
//Validates if the user has logged in before running the command
func RequiredAuthentication() {
	stg := setting.NewToolStg()
	token := stg.Token
	user := stg.User

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
