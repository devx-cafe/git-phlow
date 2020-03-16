package cmd

import (
	"fmt"
	"os"

	"github.com/code-cafe/git-phlow/cmd/cmdperm"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

// upNextCmd represents the upnext command
var upNextCmd = &cobra.Command{
	Use:   "upnext",
	Short: "get the chronologically next branch",
	Long: fmt.Sprintf(`
%s gets the next branch for integration based on its creation time.
The oldest branch is the branch that's up next.
When no --prefix flag is set, the default prefix is 'ready/'.
`, ui.Format.Bold("upnext")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		next := phlow.UpNext(options.GlobalFlagPrefixForReady)
		fmt.Fprint(os.Stdout, next)
	},
}

func init() {
	//Adding UpNext as sub-command to agentCmd
	agentCmd.AddCommand(upNextCmd)

	upNextCmd.Flags().BoolVar(&options.GlobalFlagHumanReadable, "human", false, "output human readable info")

	upNextCmd.Flags().StringVarP(&options.GlobalFlagPrefixForReady, "prefix", "p", "", "branch prefix to look for")
}
