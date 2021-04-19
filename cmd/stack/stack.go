package stack

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/pkg/stack"
)

// st holds validated Stack.
var st = stack.New()

// NewCmdStack returns a new cobra command.
func NewCmdStack() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack",
		Short: "Manage stacks",
		Long:  "Manage Grafana Cloud stacks.",
		Run:   run,
	}

	cmd.AddCommand(NewCmdCreate())
	cmd.AddCommand(NewCmdDelete())
	cmd.AddCommand(NewCmdList())

	return cmd
}

// run runs the command's action.
func run(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("failed to get help text: %s", err)
	}
}
