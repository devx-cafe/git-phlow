package cmd

import (
	//"fmt"

	"fmt"

	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "commands for automated services",
	Long: fmt.Sprintf(`
%s commands are a set of commands available for ci/cd services. They are meant as a way of services to easily be compatible with %s.

`, ui.Format.Bold("agent"), ui.Format.Bold("the praqmatic workflow")),
}

func init() {
	RootCmd.AddCommand(agentCmd)
}
