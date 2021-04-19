package grafana

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
)

// NewCmdStackApiKeyCreate returns a new cobra command.
func NewCmdApiKeyCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create STACK_SLUG NAME ROLE",
		Aliases: []string{"add"},
		Short:   "Create stack API key",
		Long:    "Create Grafana API key in the Grafana Cloud and print it out.",
		Args:    checkApiKeyCreateArgs,
		Run:     runApiKeyCreate,
	}

	cmd.Flags().BoolP("raw", "r", false, "show raw API response")
	cmd.Flags().Uint64P("seconds-to-live", "s", 0, "lifespan of the API key in seconds")

	return cmd
}

// checkStackApiKeyCreateArgs checks if the positional arguments have correct
// value. If no args are specified, it prints out the command usage.
func checkApiKeyCreateArgs(cmd *cobra.Command, args []string) error {
	argsLen := len(args)

	if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	} else if argsLen < 3 {
		return errors.New("requires STACK_SLUG, NAME and ROLE arguments")
	}

	if err := gr.SetStackSlug(args[0]); err != nil {
		return err
	}

	if err := gr.SetName(args[1]); err != nil {
		return err
	}

	if err := gr.SetRole(args[2]); err != nil {
		return err
	}

	stlFlag, err := cmd.Flags().GetUint64("seconds-to-live")
	if err != nil {
		log.Fatalf("failed to get seconds-to-live flag value: %s", err)
	}

	if stlFlag > 0 {
		if err := gr.SetSecondsToLive(stlFlag); err != nil {
			return err
		}
	}

	if token, err := common.GetToken(cmd); err == nil {
		gr.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runApiKeyCreate runs the command's action.
func runApiKeyCreate(cmd *cobra.Command, args []string) {
	key, raw, err := gr.ApiKeyCreate()
	if err != nil {
		log.Errorln("failed to create API key")
		log.Fatalln(err)
	}

	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	} else {
		fmt.Println(key)
	}
}
