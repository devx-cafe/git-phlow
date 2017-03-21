package cmd

import (
	"fmt"

	"os"
	"strconv"

	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"

	"github.com/praqma/git-phlow/cmdcheck"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web [issue]",
	Short: "open your issues on github",
	Long: `
web opens a github issue or githubs issue list based on it's arguments'
If no argument is passed, it tries to locate an issue on the current checked out branch,
if that fails it will simply open GitHubs issue list in your default browser
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		checks.IsRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if val, err := strconv.Atoi(args[0]); err == nil {
				phlow.Web(val)
			} else {
				fmt.Println("argument must be a number")
				os.Exit(0)
			}
		} else {
			phlow.Web(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(webCmd)
}
