package apikey

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
		Use:     "create ORG_SLUG NAME ROLE",
		Aliases: []string{"add"},
		Short:   "Create API keys",
		Long:    "Create Grafana Cloud API keys and print it out.",
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
	} else if argsLen < 3 {
		return fmt.Errorf("requires ORG_SLUG, NAME and ROLE argument")
	}

	if err := ak.SetOrgSlug(args[0]); err != nil {
		return err
	}

	if err := ak.SetName(args[1]); err != nil {
		return err
	}

	if err := ak.SetRole(args[2]); err != nil {
		return err
	}

	if token, err := common.GetToken(cmd); err == nil {
		ak.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runCreate runs the command's action.
func runCreate(cmd *cobra.Command, args []string) {
	key, raw, err := ak.Create()
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
