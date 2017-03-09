package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/options"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version of git phlow",
	Long: `
Get the current installed version of git-phlow including the git commit hash and date of build
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git phlow specification")
		fmt.Println("sha \t\t: " + options.Sha1)
		fmt.Println("version \t: " + options.Version)
		fmt.Println("date \t\t: " + options.Date)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
