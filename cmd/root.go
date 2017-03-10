package cmd

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/options"
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
	Run: func(cmd *cobra.Command, args []string) {
		if options.GlobalFlagVersion != false {
			fmt.Println("phlow specification")
			fmt.Println("sha \t\t: " + options.Sha1)
			fmt.Println("version \t: " + options.Version)
			fmt.Println("date \t\t: " + options.Date)
			os.Exit(0)
		}

		fmt.Println(cmd.UsageString())

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
	RootCmd.Flags().BoolVarP(&options.GlobalFlagVersion, "version", "", false, "current version")
}
