package apikey

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/pkg/grafana/apikey"
)

// ak holds validated apiKey.
var ak = apikey.New()

// NewCmdApiKey returns a new cobra command.
func NewCmdApiKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "Manage API keys",
		Long:  "Manage Grafana API keys in a Grafana Cloud stack.",
		Run:   run,
	}

	cmd.AddCommand(NewCmdCreate())

	return cmd
}

// apiKeyRun runs the command's action.
func run(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
