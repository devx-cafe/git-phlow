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
%s is required to enable 'workon' with issue management.
Auth will prompt you for a GitHub username and password to generate a token.
Don't worry, the token does not include admininstrator rights, only access to manage issues for public repositories.
`, ui.Format.Bold("auth")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.AuthCaller()
	},
}

func init() {
	RootCmd.AddCommand(authCmd)
}
