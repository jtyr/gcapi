package stack

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/common"
	"github.com/jtyr/gcapi/pkg/stack"
)

// NewCmdStackList returns a new cobra command.
func NewCmdStackList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list ORG_SLUG [STACK_SLUG]",
		Aliases: []string{"ls"},
		Short:   "List stacks",
		Long:    "List Grafana Cloud stacks.",
		Args:    checkListArgs,
		Run:     stackListRun,
	}

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

	if err := st.SetOrgSlug(args[0]); err != nil {
		return err
	}

	if argsLen == 2 {
		if err := st.SetStackSlug(args[1]); err != nil {
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

// stackListRun runs the command's action.
func stackListRun(cmd *cobra.Command, args []string) {
	list, raw, err := st.List()
	if err != nil {
		log.Errorln("failed to list stacks")
		log.Fatalln(err)
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
			if listLen > 1 {
				fmt.Printf("### %d\n", i+1)
			}

			printStackItem(&k)

			if i < listLen-1 {
				fmt.Println("")
			}
		}
	}
}

// printStackItem prints out single Stack list item.
func printStackItem(data *stack.ListItem) {
	fmt.Printf("Slug: %s\n", data.Slug)
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Println("Alert Manager Generator:")
	fmt.Printf("  - ID:  %d\n", data.AlertManagerID)
	fmt.Printf("  - URL: %s\n", data.AlertManagerGeneratorURL)
	fmt.Println("Grafana:")
	fmt.Printf("  - URL: %s\n", data.GrafanaURL)
	fmt.Println("Graphite:")
	fmt.Printf("  - ID:  %d\n", data.GraphiteID)
	fmt.Printf("  - URL: %s\n", data.GraphiteURL)
	fmt.Println("Logs:")
	fmt.Printf("  - ID:  %d\n", data.LogsID)
	fmt.Printf("  - URL: %s\n", data.LogsURL)
	fmt.Println("Prometheus:")
	fmt.Printf("  - ID:  %d\n", data.PrometheusID)
	fmt.Printf("  - URL: %s\n", data.PrometheusURL)
	fmt.Println("Traces:")
	fmt.Printf("  - ID:  %d\n", data.TracesID)
}
