package stack

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCmdApiKey returns a new cobra command.
func NewCmdStackApiKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "Manage stack API keys",
		Long:  "Manage Grafana Cloud stack API keys.",
		Run:   stackApiKeyRun,
	}

	cmd.AddCommand(NewCmdStackApiKeyCreate())

	return cmd
}

// apiKeyRun runs the command's action.
func stackApiKeyRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
