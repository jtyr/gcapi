package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GetToken returns the API authorization token defined by api-token or
// api-token-file commandline flag or by the GRAFANA_CLOUD_API_TOKEN
// environment variable.
func GetToken(cmd *cobra.Command) (string, error) {
	api_token, err := cmd.Flags().GetString("api-token")
	if err != nil {
		return "", fmt.Errorf("failed to get api-token: %s", err)
	}

	if api_token != "" {
		return api_token, nil
	}

	api_token_file, err := cmd.Flags().GetString("api-token-file")
	if err != nil {
		return "", fmt.Errorf("failed to get api-token-file: %s", err)
	}

	if api_token_file != "" {
		content, err := ioutil.ReadFile(api_token_file)
		if err != nil {
			return "", fmt.Errorf("failed to read api-token-file content: %s", err)
		}

		if len(content) == 0 {
			return "", errors.New("no token in the api-token-file")
		}

		return strings.TrimSpace(string(content)), nil
	}

	env_token := os.Getenv("GRAFANA_CLOUD_API_TOKEN")

	if env_token != "" {
		return env_token, nil
	}

	return "", fmt.Errorf("no token found")
}
