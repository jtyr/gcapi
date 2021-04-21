package apikey

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
	"github.com/jtyr/gcapi/pkg/apikey"
)

// NewCmdList returns a new cobra command.
func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list ORG_SLUG [NAME]",
		Aliases: []string{"ls"},
		Short:   "List API keys",
		Long:    "List Grafana Cloud API keys.",
		Args:    checkListArgs,
		Run:     runList,
	}

	cmd.Flags().BoolP("only-role-admin", "a", false, "show only API keys with Admin role")
	cmd.Flags().BoolP("only-role-editor", "e", false, "show only API keys with Editor role")
	cmd.Flags().BoolP("only-role-metrics-publisher", "m", false, "show only API keys with MetricsPublisher role")
	cmd.Flags().BoolP("only-role-viewer", "v", false, "show only API keys with Viewer role")
	cmd.Flags().BoolP("raw", "r", false, "show raw API response")

	return cmd
}

// checkListArgs checks if the positional arguments have correct value. If no
// args are specified, it prints out the command usage.
func checkListArgs(cmd *cobra.Command, args []string) error {
	argsLen := len(args)

	if argsLen == 0 {
		cmd.Usage()
		os.Exit(0)
	}

	if err := ak.SetOrgSlug(args[0]); err != nil {
		return err
	}

	if argsLen == 2 {
		if err := ak.SetName(args[1]); err != nil {
			return err
		}
	}

	if token, err := common.GetToken(cmd); err == nil {
		ak.SetToken(token)
	} else {
		return fmt.Errorf("failed to get authorization token: %s", err)
	}

	return nil
}

// runList runs the command's action.
func runList(cmd *cobra.Command, args []string) {
	list, raw, ec, err := ak.List()
	if err != nil {
		log.Errorf("failed to list API keys: %s", err)
		log.Exit(ec)
	}

	oraFlag, err := cmd.Flags().GetBool("only-role-admin")
	if err != nil {
		log.Fatalf("failed to get only-role-admin flag value: %s", err)
	}
	oreFlag, err := cmd.Flags().GetBool("only-role-editor")
	if err != nil {
		log.Fatalf("failed to get only-role-editor flag value: %s", err)
	}
	ormpFlag, err := cmd.Flags().GetBool("only-role-metrics-publisher")
	if err != nil {
		log.Fatalf("failed to get only-role-metrics-publisher flag value: %s", err)
	}
	orvFlag, err := cmd.Flags().GetBool("only-role-viewer")
	if err != nil {
		log.Fatalf("failed to get only-role-viewer flag value: %s", err)
	}
	rawFlag, err := cmd.Flags().GetBool("raw")
	if err != nil {
		log.Fatalf("failed to get raw flag value: %s", err)
	}

	if rawFlag {
		fmt.Println(raw)
	} else {
		listLen := len(*list)

		for i, k := range *list {
			if !(oraFlag || oreFlag || ormpFlag || orvFlag) ||
				(oraFlag && k.Role == apikey.RoleAdmin) ||
				(oreFlag && k.Role == apikey.RoleEditor) ||
				(ormpFlag && k.Role == apikey.RoleMetricsPublisher) ||
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
}

// printItem prints out single API Key list item.
func printItem(data *apikey.ListItem) {
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Role: %s\n", data.Role)
}
