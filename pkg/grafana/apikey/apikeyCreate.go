package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// apiKeyResp described properties of the API key returned by the API.
type apiKeyResp struct {
	Key string `json:"key"`
}

// Create creates a new Grafana API key and returns the value of the newly
// created API key and the raw API response.
func (ak *apiKey) Create() (string, string, error) {
	client, err := _client.New(ClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(ak.Endpoint, ak.StackSlug)

	var data []_client.Data
	data = append(data, _client.Data{Key: "name", Value: ak.Name})
	data = append(data, _client.Data{Key: "role", Value: ak.Role})

	if ak.SecondsToLive != "" {
		data = append(data, _client.Data{Key: "secondsToLive", Value: ak.SecondsToLive})
	}

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return "", "", errors.New("API key with this name already exists")
		} else {
			return "", "", err
		}
	}

	var jsonData apiKeyResp
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return jsonData.Key, string(body), nil
}
