package cmd

import (
	"fmt"

	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

//enable command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "authenticate with different services",
	Long: fmt.Sprintf(`
%s authorizes git-phlow towards a chosen issue management system like GitHub Issues and Jira. 
The auth command uses the .gitconfig file to figure out which service it tries to authorizes against. Support issue management systems right now are: 
:: Jira 
:: GitHub
`, ui.Format.Bold("auth")),
	Run: func(cmd *cobra.Command, args []string) {

		phlow.AuthCaller()

	},
}

func init() {
	RootCmd.AddCommand(authCmd)

	authCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")
}
