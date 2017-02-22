package cmd

import (
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

var (
	local          bool
	verboseDeliver bool
)

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "push changes to ready/branch",
	Long:  `deliver your work by pushing your branch to your remote, prefixed with ready`,
	Run: func(cmd *cobra.Command, args []string) {

		if local {
			phlow.LocalDeliver()
		}

	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	//Flag for local deliver
	deliverCmd.Flags().BoolVarP(&local, "local", "l", false, "local delivery")
	//Flag for verbose output
	deliverCmd.Flags().BoolVarP(&verboseDeliver, "verbose", "v", false, "verbose output")
}
