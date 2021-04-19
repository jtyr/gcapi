package stack

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
)

// NewCmdDelete returns a new cobra command.
func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete STACK_SLUG",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete stack",
		Long:    "Delete Grafana Cloud stack.",
		Args:    checkDeleteArgs,
		Run:     runDelete,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")

	return cmd
}

// checkDeleteArgs checks if the positional arguments have correct value. If no
// args are specified, it prints out the command usage.
func checkDeleteArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Usage()
		os.Exit(0)
	}

	if err := st.SetStackSlug(args[0]); err != nil {
		return err
	}

	if token, err := common.GetToken(cmd); err == nil {
		st.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runDelete runs the command's action.
func runDelete(cmd *cobra.Command, args []string) {
	raw, err := st.Delete()
	if err != nil {
		log.Errorln("failed to delete stack")
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
