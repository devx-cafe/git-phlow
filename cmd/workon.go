package cmd

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "workon [issue]",
	Short: "create a new workspace from an issue",
	Long: fmt.Sprintf(`
%s is the most essential command of git phlow. 
The command will create a new local branch from a remote integration branch, and base the name on an issue from the chosen issue management system.
If you run workon with unstaged changes in your workspace those changes will follow to the new branch. 
If you have just committed changes you will experience that those changes will disappear. This is due to the fact that when the branch is based on the remote integration branch, only changes integrated into the the remote integration branch will be available. Deliver the commits with the 'deliver' to have them availbe when you do workon. 
`, ui.Format.Bold("workon")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Fprintln(os.Stdout, cmd.UsageString())
			os.Exit(0)
		}

		phlow.WorkOnCaller(args[0])

	},
}

func init() {
	RootCmd.AddCommand(workonCmd)

	//Target for configuration
	workonCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in .gitconfig")

}
