package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/githandler"
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

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-phlow.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags()
}
