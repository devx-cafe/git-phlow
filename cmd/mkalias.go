package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// mkaliasCmd represents the mkalias command
var mkaliasCmd = &cobra.Command{
	Use:   "mkalias",
	Short: "create alias for phlow commands",
	Long: fmt.Sprintf(`
%s creates all the alias for your git phlow commands, so you can type 'git workon' in stead of
'git phlow workon', the alias will be created in your global .gitconfig file
`, options.Bold("mkalias")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.MkAlias()
	},
}

func init() {
	RootCmd.AddCommand(mkaliasCmd)
}
