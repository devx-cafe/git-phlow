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
	Short: "list issues from Task management system",
	Long: fmt.Sprintf(`
%s lists the first 30 issues from the connected issue management system. 
The are not picked in any specific order, and for Jira they are not picked by any specific project. 
The command is helpful for getting an overview of the next issues without leaving your terminal. The issue command uses the configuration 'issue-api' key-value pair to fetch the issues.
`, ui.Format.Bold("issues")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.IssueCaller()
	},
}

func init() {
	RootCmd.AddCommand(issueCmd)

	issueCmd.Flags().BoolVarP(&options.GlobalFlagMine, "mine", "m", false, "only list issues assigned to you")

	issueCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")

}
