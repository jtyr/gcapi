package apikey

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
		Use:     "delete ORG_SLUG NAME",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete API keys",
		Long:    "Delete Grafana Cloud API keys.",
		Args:    checkDeleteArgs,
		Run:     runDelete,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")

	return cmd
}

// checkDeleteArgs checks if the positional arguments have correct value. If no
// args are specified, it prints out the command usage.
func checkDeleteArgs(cmd *cobra.Command, args []string) error {
	argsLen := len(args)

	if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	} else if argsLen < 2 {
		return fmt.Errorf("requires ORG_SLUG and NAME argument")
	}

	if err := ak.SetOrgSlug(args[0]); err != nil {
		return err
	}

	if err := ak.SetName(args[1]); err != nil {
		return err
	}

	if token, err := common.GetToken(cmd); err == nil {
		ak.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runDelete runs the command's action.
func runDelete(cmd *cobra.Command, args []string) {
	raw, ec, err := ak.Delete()
	if err != nil {
		log.Errorf("failed to delete API keys: %s", err)
		log.Exit(ec)
	}

	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	}
}
