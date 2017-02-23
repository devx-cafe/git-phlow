package cmd

import (
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add work to index and commit with close# message",
	Long: `
Add all your work from working area to index, and makes an automatic close# commit, which closes
an issue on github
`,

	Run: func(cmd *cobra.Command, args []string) {
		phlow.WrapUp()
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	//Enable verbose printers
	wrapupCmd.Flags().BoolVarP(&options.GlobalFlagVerbose, "verbose", "v", false, "print verbose output")

}
