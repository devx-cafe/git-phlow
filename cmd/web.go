package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/code-cafe/git-phlow/cmd/cmdperm"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web <issue>",
	Short: "open your issues in the browser",
	Long: fmt.Sprintf(`
%s will open the default browser on the issue webpage.
It does so extracts the issue id from the current git branch, and uses the issue-web url of the configuration.
If no issue id or Jira key can be extracted, another behavior will be triggered. On Jira the dashboard will be opened. On Github the repository main page will be opened.
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
