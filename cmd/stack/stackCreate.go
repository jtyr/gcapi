package stack

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
)

// NewCmdCreate returns a new cobra command.
func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create STACK_SLUG [NAME]",
		Aliases: []string{"add"},
		Short:   "Create stack",
		Long:    "Create Grafana Cloud stack.",
		Args:    checkCreateArgs,
		Run:     runCreate,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")

	return cmd
}

// checkCreateArgs checks if the positional arguments have correct value. If no
// args are specified, it prints out the command usage.
func checkCreateArgs(cmd *cobra.Command, args []string) error {
	argsLen := len(args)

	if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	}

	if err := st.SetStackSlug(args[0]); err != nil {
		return err
	}

	if argsLen == 2 {
		if err := st.SetName(args[1]); err != nil {
			return err
		}
	} else {
		// If no Name defined, set Name to be the same like Stack Slug
		if err := st.SetName(args[0]); err != nil {
			return err
		}
	}

	if token, err := common.GetToken(cmd); err == nil {
		st.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runCreate runs the command's action.
func runCreate(cmd *cobra.Command, args []string) {
	data, raw, err := st.Create()
	if err != nil {
		log.Fatalf("failed to create stack: %s", err)
	}

	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	} else {
		printItem(data)
	}
}
