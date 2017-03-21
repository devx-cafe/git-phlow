package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmdcheck"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var cleanCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "cleanup removes all delivered branches",
	Long: fmt.Sprintf(`
%s removes all branches prefixed with 'delivered/'.
It deletes safely by running 'git branch -d'. By default, both local and remote branches are deleted.
`, ui.Bold("phlow clean")),
	PreRun: func(cmd *cobra.Command, args []string) {
		checks.IsRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		defaultBranch, _ := plugins.GitHub.Branch.Default()
		remote := githandler.ConfigBranchRemote(defaultBranch)

		phlow.Clean(remote)
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)

	//Run clean on local branches
	cleanCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "remove local delivered branches only")

	//Run clean forcefully
	cleanCmd.Flags().BoolVarP(&options.GlobalFlagForce, "force", "f", false, "force remove delivered branches")

}
