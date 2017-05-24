package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// parkCmd represents the park command
var parkCmd = &cobra.Command{
	Use:   "park",
	Short: "park the branch on the remote repo for sharing or storing",
	Long: fmt.Sprintf(`
%s pushes the current branch to the remote repository and prefixes it with 'share/'.
This allows others to 'workon' your branch, think of it as a 'remote stash' command.
To follow the phlow, don't work on it simultaneously, either pair program or let the other person take over.
`, ui.Format.Bold("park")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("park called")
	},
}

func init() {
	//RootCmd.AddCommand(parkCmd)
}
