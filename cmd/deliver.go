package cmd

import (
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "push changes to ready/branch",
	Long:  `deliver your work by pushing your branch to your remote, prefixed with ready`,
	Run: func(cmd *cobra.Command, args []string) {

		if options.GlobalFlagLocal {
			phlow.LocalDeliver()
		} else {
			phlow.Deliver()
		}

	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	//Flag for local deliver
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "local delivery")
	//Flag for verbose output
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagVerbose, "verbose", "v", false, "verbose output")
}
