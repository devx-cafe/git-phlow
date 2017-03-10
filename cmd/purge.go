package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge all delivered branches",
	Long:  `Force removes all delivered pranches and priunes origin`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("purge called")
	},
}

func init() {
	//RootCmd.AddCommand(purgeCmd)

}
