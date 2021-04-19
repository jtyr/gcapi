package apikey

import (
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Delete deletes Grafana API keys and returns the raw API response.
func (a *APIKey) Delete() (string, error) {
	// Use Grafana API token
	grafanaClientConfig := a.ClientConfig
	grafanaClientConfig.Token = a.GrafanaToken

	if a.BaseURL == "" {
		// Get Grafana API URL
		var err error
		grafanaClientConfig.BaseURL, err = a.GetGrafanaAPIURL()
		if err != nil {
			return "", fmt.Errorf("failed to get Grafana API URL: %s", err)
		}
	} else {
		grafanaClientConfig.BaseURL = a.BaseURL
	}

	client, err := _client.New(grafanaClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	list, _, err := a.List()
	if err != nil {
		return "", fmt.Errorf("failed to get API key ID: %s", err)
	}

	keyID := -1
	for _, item := range *list {
		if a.Name == item.Name {
			keyID = item.ID

			break
		}
	}

	if keyID == -1 {
		return "", errors.New("API key not found")
	}

	client.Endpoint = fmt.Sprintf(a.GrafanaEndpoint+"/%d", keyID)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", errors.New("API key not found")
		}

		return "", err
	}

	return string(body), nil
}
