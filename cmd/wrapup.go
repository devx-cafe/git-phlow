package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add work to index and commit with close# message",
	Long: `Add all your work from working area to index, and makes an automatic close# commit, which closes
	an issue on github`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("wrapup called")
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wrapupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wrapupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
