package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "bootstrap or show configuration",
	Long: fmt.Sprintf(`
%s config is the top-level command for showing and creating git-phlow .gitconfig files.
Manipulating the individual key-value pairs is done by using git's build in 'git config'.

When using a local '.gitconfig' file, you need to tell git that it should also look for configuration in the local workspace folder by adding: 
'git config --local include.path ../.gitconfig'

[phlow] is the default ini block

[phlow]
    integration_branch = master
    remote = origin
    service = jira
    issue_url = https://my.jira.instance.com
    delivery_branch_prefix = ready
`, ui.Format.Bold("config")),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

}
