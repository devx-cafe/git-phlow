package cmd

import (
	"fmt"

	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows configuration",
	Long: fmt.Sprintf(`
%s shows the configuration of the specified INI block.
If no arguments are passed it will show the default configuration git phlow will use if no '-t' flag is set when a command is executed.
This is good to use when debugging. Look for the 'scope: internal', this means that git phlow is using the internal default block, and not one found in a custom .gitconfig file.
	`, ui.Format.Bold("show")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.Show(args)
	},
}

func init() {
	configCmd.AddCommand(showCmd)

}
