package cmdperm

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/ui"
)

//RequiredAuthentication ...
//Validates if the user has logged in before running the command
func RequiredAuthentication() {
	token := githandler.ConfigGet("token", "phlow")
	user := githandler.ConfigGet("user", "phlow")

	if token == "" || user == "" {
		fmt.Printf("Please run %s to connect to github \n", ui.Format("auth").Bold)
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
