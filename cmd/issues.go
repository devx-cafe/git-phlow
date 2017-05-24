package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issues",
	Short: "list GitHub issues",
	Long: fmt.Sprintf(`
%s lists all the open issues on GitHub with its ID.
`, ui.Format.Bold("issues")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
		cmdperm.RequiredAuthentication()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.IssueList()
	},
}

func init() {
	RootCmd.AddCommand(issueCmd)

	issueCmd.Flags().BoolVarP(&options.GlobalFlagMine, "mine", "m", false, "only list issues assigned to you")
}
