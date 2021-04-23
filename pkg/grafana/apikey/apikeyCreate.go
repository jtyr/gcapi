package apikey

import (
	"encoding/json"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// createDocument describes the request JSON structure
type createDocument struct {
	Name          string `json:"name"`
	Role          string `json:"role"`
	SecondsToLive uint64 `json:"secondsToLive"`
}

// createResp described properties of the API key returned by the API.
type createResp struct {
	Key string `json:"key"`
}

// Create creates a new Grafana API key and returns the value of the newly
// created API key and the raw API response.
func (a *APIKey) Create() (string, string, error) {
	// Use Grafana API token
	grafanaClientConfig := a.ClientConfig
	grafanaClientConfig.Token = a.GrafanaToken

	if a.BaseURL == "" {
		// Get Grafana API URL
		var err error
		grafanaClientConfig.BaseURL, err = a.GetGrafanaAPIURL()
		if err != nil {
			return "", "", fmt.Errorf("failed to get Grafana API URL: %s", err)
		}
	} else {
		grafanaClientConfig.BaseURL = a.BaseURL
	}

	client, err := _client.New(grafanaClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	if a.StackSlug == "" {
		client.Endpoint = fmt.Sprintf(a.GrafanaEndpoint)
	} else {
		client.Endpoint = fmt.Sprintf(a.Endpoint, a.StackSlug)
	}

	data := createDocument{
		Name: a.Name,
		Role: a.Role,
	}

	if a.SecondsToLive != 0 {
		data.SecondsToLive = a.SecondsToLive
	}

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return "", "", fmt.Errorf("API key with this name already exists: %s", err)
		}

		return "", "", err
	}

	var jsonData createResp
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", "", fmt.Errorf("cannot parse API response as JSON: %s", err)
	}

	return jsonData.Key, string(body), nil
}
