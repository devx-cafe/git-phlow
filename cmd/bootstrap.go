package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "creates a .gitconfig file",
	Long: fmt.Sprintf(`
%s creates a new .gitconfig file in the current directory, with the default values of a git-phlow configuration. 
This is a good way to get started with the configuration file. 
	`, ui.Format.Bold("bootstrap")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.Bootstrap()
	},
}

func init() {
	configCmd.AddCommand(bootstrapCmd)

}
