package cmd

import (
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

var mkAliasVerbose bool

// mkaliasCmd represents the mkalias command
var mkaliasCmd = &cobra.Command{
	Use:   "mkalias",
	Short: "create alias for phlow commands",
	Long: `
Create all the alias for your git phlow commands, so you can type 'git workon' in stead of
'git phlow workon'`,
	Run: func(cmd *cobra.Command, args []string) {
		phlow.MkAlias(mkAliasVerbose)
	},
}

func init() {
	RootCmd.AddCommand(mkaliasCmd)

	//Verbose output
	mkaliasCmd.PersistentFlags().BoolVarP(&mkAliasVerbose, "verbose", "v", false, "verbose printers")
}
