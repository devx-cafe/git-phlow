package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

//enable command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "authenticate with github",
	Long: fmt.Sprintf(`
%s is needed to enable 'workon' with issues management.
Auth will prompt your for a github username and password,
which will generate a token we use for your github account.
Don't worry the token does not create admin access,
only acces to manage issues for public repositories'`, ui.Format.Bold("auth")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.Auth()
	},
}

func init() {
	RootCmd.AddCommand(authCmd)
}
