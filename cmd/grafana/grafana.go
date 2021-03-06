package grafana

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/grafana/apikey"
	"github.com/jtyr/gcapi/pkg/grafana"
)

// gr holds validated Grafana.
var gr = grafana.New()

// NewCmdGrafana returns a new cobra command.
func NewCmdGrafana() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grafana",
		Short: "Manage Grafana instance",
		Long:  "Manage Grafana instance.",
		Run:   grafanaRun,
	}

	cmd.AddCommand(apikey.NewCmdAPIKey())
	cmd.AddCommand(NewCmdRestart())

	return cmd
}

// grafanaRun runs the command's action.
func grafanaRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("failed to get help text: %s", err)
	}
}
