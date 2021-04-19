package grafana

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
)

// NewCmdStackList returns a new cobra command.
func NewCmdRestart() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "restart STACK_SLUG",
		Aliases: []string{"ls"},
		Short:   "Restart stack",
		Long:    "Restart Grafana Cloud stack.",
		Args:    checkRestartArgs,
		Run:     runRestart,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")

	return cmd
}

// checkRestartArgs checks if the positional arguments have correct value. If no
// args are specified, it prints out the command usage.
func checkRestartArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Usage()
		os.Exit(0)
	}

	if err := gr.SetStackSlug(args[0]); err != nil {
		return err
	}

	if token, err := common.GetToken(cmd); err == nil {
		gr.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runRestart runs the command's action.
func runRestart(cmd *cobra.Command, args []string) {
	raw, err := gr.Restart()
	if err != nil {
		log.Errorln("failed to restart stack")
		log.Fatalln(err)
	}

	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	}
}
