package grafana

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCmdApiKey returns a new cobra command.
func NewCmdApiKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "Manage API keys",
		Long:  "Manage Grafana API keys in a Grafana Cloud stack.",
		Run:   apiKeyRun,
	}

	cmd.AddCommand(NewCmdApiKeyCreate())

	return cmd
}

// apiKeyRun runs the command's action.
func apiKeyRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
