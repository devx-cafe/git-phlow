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
	Use:   "deliver",
	Short: "deliver work in agent version",
	Long: fmt.Sprintf(`
%s delivers the changes in an agent version. The current branch if will be merged into the default branch
with your default branch and pushed to the remote.
`, ui.Format("integrate").Bold),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, err := plugins.GitHub.Branch.Default()
		if err != nil {
			fmt.Println(err)
		}
		phlow.LocalDeliver(defaultBranch)
	},
}

func init() {
	agentCmd.AddCommand(integrateCmd)
}
