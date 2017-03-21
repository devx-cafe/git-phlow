package checks

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/ui"
)

//IsAuthenticated ...
//Validates if the user has logged in before running the command
func IsAuthenticated() {
	token := githandler.ConfigGet("token", "phlow")
	user := githandler.ConfigGet("user", "phlow")

	if token == "" || user == "" {
		fmt.Printf("Please run %s to connect to github \n", ui.Bold("auth"))
		os.Exit(0)
	}
}

//IsRepository ...
//Validates if the command is runnign in a git repository
func IsRepository() {
	if err := githandler.Status(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(0)
	}
}
