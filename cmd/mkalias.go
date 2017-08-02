package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/phlow"
)

// mkaliasCmd represents the mkalias command
var mkaliasCmd = &cobra.Command{
	Use:   "mkalias",
	Short: "create aliases for phlow commands",
	Long: fmt.Sprintf(`
%s creates aliases for the git phlow commands.
This allows you to use 'git workon', rather than 'git phlow workon'.
The aliases are added to your global .gitconfig file.
`, ui.Format.Bold("mkalias")),
	Run: func(cmd *cobra.Command, args []string) {
		phlow.MakeAliasCaller()
	},
}

func init() {
	RootCmd.AddCommand(mkaliasCmd)
}
