package cmd

import (
	"fmt"
	"os"
	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "workon [issue number]",
	Short: "create or change to an issue branch",
	Long: fmt.Sprintf(`
%s creates a new branch from an issue number fetched from GitHub.
A new branch will be created, based on your remote default branch and named after the issue title, e.g. "42-calculate-meaning-of-life".
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
	workonCmd.Flags().StringVarP(&options.GlobalFlagTarget, "target", "t", "", "the name of the INI block in your .phlow files")

}
