package apikey

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
	"github.com/jtyr/gcapi/pkg/consts"
	"github.com/jtyr/gcapi/pkg/grafana/apikey"
)

// NewCmdList returns a new cobra command.
func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list (ORG_SLUG STACK_SLUG|--grafana-api-url STRING) [NAME]",
		Aliases: []string{"add"},
		Short:   "List API keys",
		Long:    "List Grafana API keys in a specific Stack of the Grafana Cloud and print them out.",
		Args:    checkListArgs,
		Run:     runList,
	}

	cmd.Flags().BoolP("only-role-admin", "a", false, "show only API keys with Admin role")
	cmd.Flags().BoolP("only-role-editor", "e", false, "show only API keys with Editor role")
	cmd.Flags().BoolP("only-role-viewer", "v", false, "show only API keys with Viewer role")
	cmd.Flags().BoolP("raw", "r", false, "show raw API response")
	cmd.Flags().StringP("grafana-api-url", "g", "", "Grafana API URL (e.g. https://grafana.domain.com/api).")

	return cmd
}

// checkListArgs checks if the positional arguments have correct
// value. If no args are specified, it prints out the command usage.
func checkListArgs(cmd *cobra.Command, args []string) error {
	gauFlag, err := common.GetGrafanaAPIURL(cmd)
	if err != nil {
		log.Fatalf("failed to get Grafana API URL: %s", err)
	}

	argsLen := len(args)

	if gauFlag != "" {
		if err := ak.SetBaseURL(gauFlag); err != nil {
			return err
		}

		if argsLen > 1 {
			return errors.New("requires only optional NAME argument")
		}

		if argsLen == 1 {
			if err := ak.SetName(args[0]); err != nil {
				return err
			}
		}
	} else if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	} else if argsLen < 2 {
		return errors.New("requires ORG_SLUG and STACK_SLUG argument")
	} else if argsLen > 2 {
		return errors.New("requires only ORG_SLUG and STACK_SLUG argument")
	} else {
		if err := ak.SetOrgSlug(args[0]); err != nil {
			return err
		}

		if err := ak.SetStackSlug(args[1]); err != nil {
			return err
		}

		if argsLen == 3 {
			if err := ak.SetName(args[2]); err != nil {
				return err
			}
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

// runList runs the command's action.
func runList(cmd *cobra.Command, args []string) {
	list, raw, err := ak.List()
	if err != nil {
		log.Fatalf("failed to get API keys: %s", err)
	}

	oraFlag, err := cmd.Flags().GetBool("only-role-admin")
	if err != nil {
		log.Fatalf("failed to get only-role-admin flag value: %s", err)
	}
	oreFlag, err := cmd.Flags().GetBool("only-role-editor")
	if err != nil {
		log.Fatalf("failed to get only-role-editor flag value: %s", err)
	}
	orvFlag, err := cmd.Flags().GetBool("only-role-viewer")
	if err != nil {
		log.Fatalf("failed to get only-role-viewer flag value: %s", err)
	}
	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	printed := false

	if rawFlag {
		fmt.Println(raw)
	} else {
		listLen := len(*list)

		for i, k := range *list {
			if ak.Name != "" {
				if k.Name == ak.Name {
					printItem(&k)
					printed = true
				}
			} else if !(oraFlag || oreFlag || orvFlag) ||
				(oraFlag && k.Role == apikey.RoleAdmin) ||
				(oreFlag && k.Role == apikey.RoleEditor) ||
				(orvFlag && k.Role == apikey.RoleViewer) {

				if listLen > 1 {
					fmt.Printf("### %d\n", i+1)
				}

				printItem(&k)

				if i < listLen-1 {
					fmt.Println("")
				}
			}
		}
	}

	if !printed && ak.Name != "" {
		log.Error("failed to list API keys: API key not found")
		log.Exit(consts.ExitNotFound)
	}
}

// printItem prints out single API Key list item.
func printItem(data *apikey.ListItem) {
	fmt.Printf("ID: %d\n", data.ID)
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Role: %s\n", data.Role)
}
