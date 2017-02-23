package cmd

import (
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

//enable command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "authenticate and enable phlow",
	Long: `
authenticate git phlow with your github account for issue management and issue branch workflow, it
requires you to enter username and  password`,
	Run: func(cmd *cobra.Command, args []string) {

		phlow.Auth()

	},
}

func init() {
	RootCmd.AddCommand(authCmd)

}
