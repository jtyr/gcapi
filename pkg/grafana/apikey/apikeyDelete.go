package apikey

import (
	"errors"
	"fmt"
	"net/http"

	_client "github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
)

var deleteClient *http.Client

// Delete deletes Grafana API keys and returns the raw API response.
func (a *APIKey) Delete() (string, int, error) {
	// Use Grafana API token
	grafanaClientConfig := a.ClientConfig
	grafanaClientConfig.Token = a.GrafanaToken

	if a.BaseURL == "" {
		// Get Grafana API URL
		var err error
		grafanaClientConfig.BaseURL, err = a.GetGrafanaAPIURL()
		if err != nil {
			return "", consts.ExitError, fmt.Errorf("failed to get Grafana API URL: %s", err)
		}
	} else {
		grafanaClientConfig.BaseURL = a.BaseURL
	}

	client, err := _client.New(grafanaClientConfig)
	if err != nil {
		return "", consts.ExitError, fmt.Errorf("failed to get client: %s", err)
	}

	list, _, err := a.List()
	if err != nil {
		return "", consts.ExitError, fmt.Errorf("failed to get API key ID: %s", err)
	}

	keyID := -1
	for _, item := range *list {
		if a.Name == item.Name {
			keyID = item.ID

			break
		}
	}

	if keyID == -1 {
		return "", consts.ExitNotFound, errors.New("API key not found in the list of API keys")
	}

	// This is here only to be able to mock the Delete() response
	if deleteClient != nil && deleteClient.Transport != nil {
		_client.Client = deleteClient
		client, _ = _client.New(grafanaClientConfig)
	}

	client.Endpoint = fmt.Sprintf(a.GrafanaEndpoint+"/%d", keyID)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", consts.ExitNotFound, fmt.Errorf("API key not found: %s", err)
		}

		return "", consts.ExitError, err
	}

	return string(body), consts.ExitOk, nil
}
