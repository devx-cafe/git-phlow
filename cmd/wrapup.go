package cmd

import (
	"fmt"

	"github.com/code-cafe/git-phlow/cmd/cmdperm"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add changes to index and auto commit",
	Long: fmt.Sprintf(`
%s is used to collect all the changes just made and put them into a commit, that is ready to be delivered into the integration branch.
The command will add changes from the workspace to the index/staging area of git, and commit with 'smart commit' syntax that will close the mentioned issue when it is integrated on the integration branch.
The smart commits only works on systems that support it.
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
