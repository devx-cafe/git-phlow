package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// integrateCmd represents the integrate command
var integrateCmd = &cobra.Command{
	Use:   "integrate",
	Short: "integrate work in agent version",
	Long: fmt.Sprintf(`
%s delivers the changes as the agent.
The current branch will be merged into the default branch and pushed to the remote.
`, ui.Format.Bold("integrate")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, err := plugins.GitHub.Default()
		if err != nil {
			fmt.Println(err)
		}
		phlow.LocalDeliver(defaultBranch)
	},
}

func init() {
	agentCmd.AddCommand(integrateCmd)
}
