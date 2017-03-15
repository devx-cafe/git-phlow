package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/ui"
)

// upNextCmd represents the upnext command
var upNextCmd = &cobra.Command{
	Use:   "upnext",
	Short: "get the chronologically next branch",
	Long: fmt.Sprintf(`
%s gets the next branch ready for integration based on the branch creation time.
The branch created first, is the branch thats up next.
`, ui.Bold("upnext")),

	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)
		remote := githandler.ConfigBranchRemote(defaultBranch)

		phlow.UpNext(remote)
	},
}

func init() {
	//Adding UpNext as sub-command to agentCmd
	agentCmd.AddCommand(upNextCmd)

	upNextCmd.Flags().BoolVar(&options.GlobalFlagHumanReadable, "human", false, "output human readable info")

}
