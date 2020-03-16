package cmd

import (
	"fmt"

	"github.com/code-cafe/git-phlow/cmd/cmdperm"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver [args]",
	Short: "deliver changes to remote master",
	Long: fmt.Sprintf(`
%s is the command used to hand over the work you just commited with 'wrapup'. 
checkout the issue branch you need to deliver and run the command. This will push local branch to the repository and prefixed it with 'ready/'. 
On the repository it is ready to be integrated automatically into the integration branch, if a CI server is configured. 

When delivered, local branches will be prefixed with 'delivered/'

using the '--local' option will integrate the branch locally and push the changes on the integration branch
`, ui.Format.Bold("deliver")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		phlow.DeliverCaller()

	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	//Flag for local deliver
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "local delivery")

	deliverCmd.Flags().BoolVarP(&options.GlobalFlagShowTestOutput, "showtest", "s", false, "show test output")

	deliverCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")
}
