package apikey

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
)

// NewCmdDelete returns a new cobra command.
func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete (ORG_SLUG STACK_SLUG|--grafana-api-url STRING) NAME",
		Aliases: []string{"add"},
		Short:   "Delete Grafana API key",
		Long:    "Delete Grafana API keys in a specific Stack of the Grafana Cloud and print them out.",
		Args:    checkDeleteArgs,
		Run:     runDelete,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")
	cmd.Flags().StringP("grafana-api-url", "g", "", "Grafana API URL (e.g. https://grafana.domain.com/api).")

	return cmd
}

// checkDeleteArgs checks if the positional arguments have correct
// value. If no args are specified, it prints out the command usage.
func checkDeleteArgs(cmd *cobra.Command, args []string) error {
	gauFlag, err := common.GetGrafanaApiURL(cmd)
	if err != nil {
		log.Fatalf("failed to get Grafana API URL: %s", err)
	}

	argsLen := len(args)

	if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	} else if gauFlag != "" {
		if argsLen == 0 {
			return errors.New("requires NAME argument")
		}

		if err := ak.SetName(args[0]); err != nil {
			return err
		}

		if err := ak.SetBaseURL(gauFlag); err != nil {
			return err
		}
	} else if argsLen < 3 {
		return errors.New("requires ORG_SLUG, STACK_SLUG and NAME argument")
	} else {
		if err := ak.SetOrgSlug(args[0]); err != nil {
			return err
		}

		if err := ak.SetStackSlug(args[1]); err != nil {
			return err
		}

		if err := ak.SetName(args[2]); err != nil {
			return err
		}

		if token, err := common.GetToken(cmd); err == nil {
			ak.SetToken(token)
		} else {
			return fmt.Errorf("failed to get authorization token: %s", err)
		}
	}

	if token, err := common.GetGrafanaToken(cmd); err == nil {
		ak.SetGrafanaToken(token)
	} else {
		return fmt.Errorf("failed to get Grafana authorization token: %s", err)
	}

	return nil
}

// runDelete runs the command's action.
func runDelete(cmd *cobra.Command, args []string) {
	raw, err := ak.Delete()
	if err != nil {
		log.Fatalf("failed to delete API key: %s", err)
	}

	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	}
}
