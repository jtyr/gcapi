package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// createResp described properties of the API key returned by the API.
type createResp struct {
	Key string `json:"key"`
}

// Create creates a new Grafana API key and returns the value of the newly
// created API key and the raw API response.
func (a *APIKey) Create() (string, string, error) {
	client, err := _client.New(a.ClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(a.Endpoint, a.StackSlug)

	var data []_client.Data
	data = append(data, _client.Data{Key: "name", Value: a.Name})
	data = append(data, _client.Data{Key: "role", Value: a.Role})

	if a.SecondsToLive != "" {
		data = append(data, _client.Data{Key: "secondsToLive", Value: a.SecondsToLive})
	}

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return "", "", errors.New("API key with this name already exists")
		}

		return "", "", err
	}

	var jsonData createResp
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return jsonData.Key, string(body), nil
}
