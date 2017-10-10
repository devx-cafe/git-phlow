package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web [issue]",
	Short: "open your issues in the browser",
	Long: fmt.Sprintf(`
%s opens an issue in your default browser

The command uses a targetet configuration to figure out where you have your issues hosted e.g. Github or jira.
Use the 'issue-web' field in the configuration to point to the url of the issue management system.

The command can take an argument of the issue you want to use.
If no argument is given, it tries to find the issue of the currently checked out branch, if that fails it simply opens the GitHub issue list in your default browser.
`, ui.Format.Bold("web")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		phlow.WebCaller(args)

	},
}

func init() {
	RootCmd.AddCommand(webCmd)

	webCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")
}
