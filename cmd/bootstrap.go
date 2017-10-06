package cmd

import (
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "creates a .gitconfig file",
	Long: `
Creates a new .gitconfig file in your local repository
	`,
	Run: func(cmd *cobra.Command, args []string) {
		phlow.Bootstrap()
	},
}

func init() {
	configCmd.AddCommand(bootstrapCmd)

}
