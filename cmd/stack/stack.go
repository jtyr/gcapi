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
		Run:   stackRun,
	}

	cmd.AddCommand(NewCmdStackCreate())
	cmd.AddCommand(NewCmdStackDelete())
	cmd.AddCommand(NewCmdStackList())

	return cmd
}

// stackRun runs the command's action.
func stackRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
