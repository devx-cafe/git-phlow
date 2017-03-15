package cmd

import (
	//"fmt"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/ui"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "commands for automated services",
	Long: fmt.Sprintf(`
%s commands are a set of subcommands available for ci/cd services. They are meant as a way of services to easily
be compatible with %s.
These are not meant to used locally as a part of your local workflow
`, ui.Bold("agent"), ui.Bold("the praqmatic workflow")),
}

func init() {
	RootCmd.AddCommand(agentCmd)
}
