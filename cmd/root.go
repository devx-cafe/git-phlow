package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/flags"
	"github.com/code-cafe/git-phlow/setting"

	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "git-phlow",
	Short: "git extension for phlow workflow",
	Long: fmt.Sprintf(`
%s is a git extension that provides an extra set of commands, enabling you to:
create, work on and deliver tasks.

USAGE: git phlow COMMAND

`, ui.Format.Bold("git-phlow")),
	Run: RootCommand,
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

	RootCmd.Flags().BoolVar(&flags.RootVersion, "version", false, "show version")

	RootCmd.Flags().BoolVar(&flags.RootInit, "init", false, "initialize a .gitconfig file with default phlow")

	RootCmd.PersistentFlags().BoolVar(&options.GlobalFlagNoColor, "no-color", false, "remove color")

	RootCmd.PersistentFlags().BoolVarP(&options.GlobalFlagVerbose, "verbose", "v", false, "verbose output")
}

// RootCommand ...
func RootCommand(cmd *cobra.Command, args []string) {

	if flags.RootVersion != false {
		rootVersion()
	}

	if flags.RootInit != false {
		rootInit(func() string {
			if path, err := executor.Run("git", "rev-parse", "--show-toplevel"); err == nil {
				return path
			}
			panic("could not get repository root")
		}())
	}

	if flags.RootInit == false && flags.RootVersion == false {
		fmt.Println(cmd.UsageString())
	}
}

func rootVersion() {
	info := make(map[string]string)
	info["sha"] = options.Sha1
	info["version"] = options.Version
	info["date"] = options.Date

	jsonBytes, err := json.Marshal(info)

	if err != nil {
		log.Fatal("version could not be shown")
	}

	fmt.Println(string(jsonBytes))

}

func rootInit(path string) {
	err := setting.BootstrapPhlowConfig(strings.TrimSpace(path), "master")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not initialize configuration")
	}
}
