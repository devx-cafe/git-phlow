package cmd

import (
	"fmt"
	"github.com/praqma/git-phlow/phlowimpl"
	"github.com/spf13/cobra"
	"os"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "work-on [issue number]",
	Short: "create a branch from an issue",
	Long: `creates a new branch named from a issue on a selected plugin`,
	Run: func(cmd *cobra.Command, args []string) {

		message, err := phlowimpl.WorkOn(args)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, message)
	},
}

func init() {
	RootCmd.AddCommand(workonCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
