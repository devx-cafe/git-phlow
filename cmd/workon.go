package cmd

import (
	"fmt"

	"os"
	"strconv"

	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "workon [issue number]",
	Short: "create or change to an issue branch",
	Long: `
[workon] creates a new branch from an issue number fetched from github.
A branch will be created from your remote default branch and named after the title
of the issue eg. branch => "42-calculate-meaning-of-life".
From here you the work can start, and when finished run phlow wrapup.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			issue, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stdout, "Whoops \nYour argument: %s, is not a 'number' I only feast on numbers \n", args[0])
				os.Exit(0)
			}

			phlow.WorkOn(issue)
			os.Exit(0)
		}

		fmt.Fprintln(os.Stdout, cmd.UsageString())
	},
}

func init() {
	RootCmd.AddCommand(workonCmd)

	//Enable verbose printers
	workonCmd.PersistentFlags().BoolVarP(&options.GlobalFlagVerbose, "verbose", "v", false, "verbose printers")

}
