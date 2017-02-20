package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Local string

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "push changes to ready/branch",
	Long: `deliver your work by pushing your branch to your remote, prefixed with ready`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("deliver called")
	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deliverCmd.PersistentFlags().String("foo", "", "A help for foo")

	RootCmd.Flags().StringVarP(&Local, "local", "l", "", "Do a local deliver")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deliverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
