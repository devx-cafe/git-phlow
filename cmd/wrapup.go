package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add changes to index and auto commit",
	Long: fmt.Sprintf(`
%s adds the files in the workin directory to the index and makes a commit.
The commit message generated from the branch name and prepends a smart commit function to close the issue.
The smart commit work with GitHub and Jira.
`, ui.Format.Bold("wrapup")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.WrapUpCaller()
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	wrapupCmd.Flags().StringVar(&options.GlobalFlagForceMessage, "force", "", "use a custom commit message instead")

	wrapupCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")
}
