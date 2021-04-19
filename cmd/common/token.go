package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GetToken returns the API authorization token defined by cloud-api-token or
// cloud-api-token-file commandline flag or by the GRAFANA_CLOUD_API_TOKEN
// environment variable.
func GetToken(cmd *cobra.Command) (string, error) {
	apiToken, err := cmd.Flags().GetString("cloud-api-token")
	if err != nil {
		return "", fmt.Errorf("failed to get cloud-api-token: %s", err)
	}

	if apiToken != "" {
		return apiToken, nil
	}

	apiTokenFile, err := cmd.Flags().GetString("cloud-api-token-file")
	if err != nil {
		return "", fmt.Errorf("failed to get cloud-api-token-file: %s", err)
	}

	if apiTokenFile != "" {
		content, err := ioutil.ReadFile(apiTokenFile)
		if err != nil {
			return "", fmt.Errorf("failed to read cloud-api-token-file content: %s", err)
		}

		if len(content) == 0 {
			return "", errors.New("no token in the cloud-api-token-file")
		}

		return strings.TrimSpace(string(content)), nil
	}

	envToken := os.Getenv("GRAFANA_CLOUD_API_TOKEN")

	if envToken != "" {
		return envToken, nil
	}

	return "", fmt.Errorf("no token found")
}

// GetGrafanaToken returns the API authorization token defined by
// grafana-api-token or grafana-api-token-file commandline flag or by the
// GRAFANA_API_TOKEN environment variable.
func GetGrafanaToken(cmd *cobra.Command) (string, error) {
	apiToken, err := cmd.Flags().GetString("grafana-api-token")
	if err != nil {
		return "", fmt.Errorf("failed to get grafana-api-token: %s", err)
	}

	if apiToken != "" {
		return apiToken, nil
	}

	apiTokenFile, err := cmd.Flags().GetString("grafana-api-token-file")
	if err != nil {
		return "", fmt.Errorf("failed to grafana-get api-token-file: %s", err)
	}

	if apiTokenFile != "" {
		content, err := ioutil.ReadFile(apiTokenFile)
		if err != nil {
			return "", fmt.Errorf("failed to read grafana-api-token-file content: %s", err)
		}

		if len(content) == 0 {
			return "", errors.New("no token in the grafana-api-token-file")
		}

		return strings.TrimSpace(string(content)), nil
	}

	envToken := os.Getenv("GRAFANA_API_TOKEN")

	if envToken != "" {
		return envToken, nil
	}

	return "", fmt.Errorf("no token found")
}
