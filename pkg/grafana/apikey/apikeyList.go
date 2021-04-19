package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// ListItem described properties of individual List item returned by the API.
type ListItem struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// listResp described properties of the document returned by the API.
type listResp []ListItem

// List lists Grafana API keys and returns the list and the raw API response.
func (a *apiKey) List() (*listResp, string, error) {
	// Use Grafana API token
	grafanaClientConfig := a.ClientConfig
	grafanaClientConfig.Token = a.GrafanaToken

	if a.BaseURL == "" {
		// Get Grafana API URL
		var err error
		grafanaClientConfig.BaseURL, err = a.GetGrafanaApiURL()
		if err != nil {
			return nil, "", fmt.Errorf("failed to get Grafana API URL: %s", err)
		}
	} else {
		grafanaClientConfig.BaseURL = a.BaseURL
	}

	client, err := _client.New(grafanaClientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = a.GrafanaEndpoint

	body, statusCode, err := client.Get()
	if err != nil {
		if statusCode == 404 {
			return nil, "", errors.New("Grafana instance not found")
		} else {
			return nil, "", err
		}
	}

	var jsonData listResp
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return &jsonData, string(body), nil
}
