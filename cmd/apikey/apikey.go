package apikey

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/pkg/apikey"
)

// ak holds validated ApiKey.
var ak = apikey.New()

// NewCmdApiKey returns a new cobra command.
func NewCmdApiKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "Manage API keys",
		Long:  "Manage Grafana Cloud API keys.",
		Run:   run,
	}

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
