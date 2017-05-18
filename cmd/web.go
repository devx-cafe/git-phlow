package cmd

import (
	"fmt"

	"os"
	"strconv"

	"github.com/praqma/git-phlow/phlow"
	"github.com/spf13/cobra"

	"github.com/praqma/git-phlow/cmd/cmdperm"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web [issue]",
	Short: "open your issues on github",
	Long: fmt.Sprintf(`
%s opens a GitHub issue or GitHub issue list based on it's arguments.
If no argument is given, it tries to find the issue of the currently checked out branch, if that fails it simply opens the GitHub issue list in your default browser.
`, ui.Format.Bold("web")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
		cmdperm.RequiredAuthentication()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if val, err := strconv.Atoi(args[0]); err == nil {
				phlow.Web(val)
			} else {
				fmt.Println("Argument must be a number")
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
