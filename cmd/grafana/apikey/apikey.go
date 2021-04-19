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
		Long:  "Manage Grafana API keys.",
		Run:   run,
	}

	cmd.PersistentFlags().StringP(
		"grafana-api-token", "T", "",
		"token used to authenticate to the Grafana API")
	cmd.PersistentFlags().StringP(
		"grafana-api-token-file", "F", "",
		"path to a file containing the token used to authenticate to the Grafana API")

	cmd.AddCommand(NewCmdCreate())
	cmd.AddCommand(NewCmdDelete())
	cmd.AddCommand(NewCmdList())

	return cmd
}

// run runs the command's action.
func run(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Errorln("failed to get help text")
		log.Fatalln(err)
	}
}
