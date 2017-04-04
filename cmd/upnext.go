package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// upNextCmd represents the upnext command
var upNextCmd = &cobra.Command{
	Use:   "upnext",
	Short: "get the chronologically next branch",
	Long: fmt.Sprintf(`
%s gets the next branch ready for integration based on the branch creation time.
The branch created first, is the branch thats up next.
if no --prefix flag is set, the default prefix is ready/
`, ui.Format.Bold("upnext")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GitHub.Default()
		remote := githandler.ConfigBranchRemote(defaultBranch)

		rdy := phlow.UpNext(remote, options.GlobalFlagPrefixForReady)
		fmt.Println(rdy)
	},
}

func init() {
	//Adding UpNext as sub-command to agentCmd
	agentCmd.AddCommand(upNextCmd)

	upNextCmd.Flags().BoolVar(&options.GlobalFlagHumanReadable, "human", false, "output human readable info")

	upNextCmd.Flags().StringVarP(&options.GlobalFlagPrefixForReady, "prefix", "p", "", "prefix branches ready for integration")
}
