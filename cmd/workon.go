package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "workon [issue number]",
	Short: "Create or change branch from an issue number",
	Long: `
	work-on creates a new branch from an issue number, if phlow have been
	initialized with a plugin, it will automatically fetch the issue name from the extension, otherwise
	it will create a default name prefixed with your issue number
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			issue, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stdout, "Whoops\n your argument: %s, is not a 'number' I only feast on numbers \n", args[0])
				os.Exit(0)
			}
			phlow.WorkOn(issue)
		}
		fmt.Fprintln(os.Stdout, cmd.UsageString())
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
