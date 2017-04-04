package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/cmd/cmdperm"
)

// parkCmd represents the park command
var parkCmd = &cobra.Command{
	Use:   "park",
	Short: "park the branch on the remote repo for sharing or storing",
	Long: `
park pushes to current branch to the remote repository and prefixes it with share/, so others can do a workon
on the branch.
Following the flow you should not work on it simultaniously; either pair program on one computer, or let the other person
take over for either reviwing the work or taking over the work.
Think of it as a remote stash command.
`,
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
