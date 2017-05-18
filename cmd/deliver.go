package cmd

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver [test args]",
	Short: "deliver changes to remote master",
	Long: fmt.Sprintf(`
%s pushes your committed changes to the remote repository.
Deliver only runs if you are located on an issue branch created with workon.
The delivered branch's name will prefixed 'ready/', so the integration services supporting the workflow can pick up your changes.

If you deliver with the 'local' option, the issue branch will be merged with your default branch, pushed to your remote's default branch and prefixed with 'delivered/'.
`, ui.Format.Bold("deliver")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GitHub.Default()
		//Run tests before deliver
		if len(args) > 0 {
			if err := phlow.TestDeliver(args); err != nil {
				fmt.Println(ui.Format.Error("!! Tests did not exit with code 0 !!"))
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(ui.Format.Success("Tests exited with code 0 - starting deliver process"))
		}

		//If Run if local deliver
		if options.GlobalFlagLocal {
			phlow.LocalDeliver(defaultBranch)
			return
		}
		//Deliver with ready branch
		phlow.Deliver(defaultBranch)
	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	//Flag for local deliver
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "local delivery")
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagShowTestOutput, "showtest", "s", false, "show test output")
}
