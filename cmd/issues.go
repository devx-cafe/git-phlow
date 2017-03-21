package cmd

import (
	"github.com/praqma/git-phlow/cmdcheck"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issues",
	Short: "list GitHub issues",
	Long: `
List all the open issues on GitHub with it's ID. Helps you locate what needs to be worked on
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		checks.IsRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.IssueList()
	},
}

func init() {
	RootCmd.AddCommand(issueCmd)

	issueCmd.Flags().BoolVarP(&options.GlobalFlagMine, "mine", "m", false, "list only issues assigned to you")
}
