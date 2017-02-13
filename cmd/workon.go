package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"fmt"
	"os"
	"github.com/praqma/git-phlow/phlowimpl"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "work-on [issue number]",
	Short: "Create or change branch from an issue number",
	Long:  `creates a new branch named from a issue on a selected plugin`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			issueNum, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stdout, "Whoops\n your argument: %s, is not a 'number' i only feast on numbers \n", args[0])
				os.Exit(0)
			}
			phlowimpl.WorkOn(issueNum)
		}
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
