package cmd

import (
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"
)

// deliverCmd represents the deliver command
var deliverCmd = &cobra.Command{
	Use:   "deliver",
	Short: "deliver changes to remote master",
	Long: `
[deliver] fetches the latest changes and tries to push the committed changes
to the remote repository. Deliver only runs if you are located on an issue branch
created with workon, or a branch which if not your default (typically master).
the delivered branch will prefix the branchname with "ready/",
so integration services supporting the workflow can pick up the changes.

    --local - if you deliver with local, the branch will be merged
	      with your default branch and pushed to your "remote default branch"
	      and prefixed with "/delivered"
`,
	Run: func(cmd *cobra.Command, args []string) {

		if options.GlobalFlagLocal {
			phlow.LocalDeliver()
		} else {
			phlow.Deliver()
		}

	},
}

func init() {
	RootCmd.AddCommand(deliverCmd)

	//Flag for local deliver
	deliverCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "local delivery")
}
