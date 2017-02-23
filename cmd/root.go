package cmd

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "git-phlow",
	Short: "git extension for phlow workflow",
	Long: `
git-phlow is a git extension that provides an extra set of commands, enabling you to:
workon, deliver and create new tasks.
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//Move as before check
		if err := githandler.Status(); err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(0)
		}

	},
}

//Execute ...
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags()
}
