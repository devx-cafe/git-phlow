package cmd

import (
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows configuration",
	Long: `
Shows a specified configuration block on your configuration file.
If no arguments are passed it will show the internal default configuration
	`,
	Run: func(cmd *cobra.Command, args []string) {
		phlow.Show(args)
	},
}

func init() {
	configCmd.AddCommand(showCmd)

}
