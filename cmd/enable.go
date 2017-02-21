package cmd

import (
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

var verboseEnable bool

//enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "authenticate and enable phlow",
	Long: `authenticate git phlow with your github account for issue management and issue branch workflow, it
	requires you to enter username and  password`,
	Run: func(cmd *cobra.Command, args []string) {

		phlow.Enable(verboseEnable)
	},
}

func init() {
	RootCmd.AddCommand(enableCmd)

	//verbose output flag
	wrapupCmd.Flags().BoolVarP(&verboseEnable, "verbose", "v", false, "enable verpose print")
}
