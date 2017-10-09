package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "bootstrap or show configuration",
	Long: `
git phlow uses gits own configuration, so all your configurations can be placed everywhere git config can get them.
Best practice is to create a .gitconfig in your repository with your won configurations.

Git phlow comes with internal defaults configured to work natively with github. So you can use git phlow
without any configuration at all. If you want to customize it the default setting create an INI-block
in a local .gitconfig, and make sure all mandatory fields are set.

To use a local .gitconfig file remember to add it to the config path, with this command
	'git config --local include.path ../.gitconfig'

[phlow] is the default ini block

Default configuration:
	[phlow]
    integration_branch = master
    remote = origin
    service = jira
    issue_url = https://my.jira.instance.com
    delivery_branch_prefix = ready
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

}
