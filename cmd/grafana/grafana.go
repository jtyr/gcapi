package grafana

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/pkg/grafana"
)

// gr holds validated Grafana.
var gr = grafana.New()

// NewCmdStack returns a new cobra command.
func NewCmdGrafana() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grafana",
		Short: "Manage Grafana",
		Long:  "Manage Grafana in Grafana Cloud.",
		Run:   grafanaRun,
	}

	cmd.AddCommand(NewCmdApiKey())
	cmd.AddCommand(NewCmdRestart())

	return cmd
}

// grafanaRun runs the command's action.
func grafanaRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
