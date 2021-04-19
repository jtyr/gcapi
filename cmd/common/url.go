package common

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GetGrafanaAPIURL returns the Grafan API URL defined by grafana-api-url
// commandline flag or by the GRAFANA_API_URL environment variable.
func GetGrafanaAPIURL(cmd *cobra.Command) (string, error) {
	apiURL, err := cmd.Flags().GetString("grafana-api-url")
	if err != nil {
		return "", fmt.Errorf("failed to get grafana-api-url: %s", err)
	}

	if apiURL != "" {
		return apiURL, nil
	}

	envURL := os.Getenv("GRAFANA_API_URL")

	if envURL != "" {
		return envURL, nil
	}

	return "", nil
}
