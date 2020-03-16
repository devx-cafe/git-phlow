package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l", "lo"},
	Example: "git phlow login login",
	Short:   "I am the short desc",
	Long:    `I am the long desc`,
	ArgAliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Facking shit")

	},
}

func init() {
	RootCmd.AddCommand(loginCommand)
	loginCommand.AddCommand(&cobra.Command{Use: "hello [hello]", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hgello world")
	}})
	loginCommand.Flags().


}
