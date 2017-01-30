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
	Short: "create a branch annotated with issue",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		message, err := phlowimpl.Workon(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
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
