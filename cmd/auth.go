package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/options"
)

//enable command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "authenticate with different services",
	Long: fmt.Sprintf(`
%s will authenticate to a service specified in your .phlow or .gitconfig files. It will use the default service if no other is specified.
Auth supports two services:
- Jira
- GitHub
`, ui.Format.Bold("auth")),
	Run: func(cmd *cobra.Command, args []string) {

		phlow.AuthCaller()

	},
}

func init() {
	RootCmd.AddCommand(authCmd)

	authCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in your .phlow files")
}
