package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// createDocument describes the request JSON structure
type createDocument struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// createResp describes the structure of the JSON document returned by the API.
type createResp struct {
	Token string `json:"token"`
}

// Create creates a new API key and returns the value of newly created API key
// and the raw API response.
func (a *APIKey) Create() (string, string, error) {
	client, err := _client.New(a.ClientConfig)
	if err != nil {
		return "", "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(a.Endpoint, a.OrgSlug)

	data := createDocument{
		Name: a.Name,
		Role: a.Role,
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
		return "", "", fmt.Errorf("cannot parse API response as JSON: %s", err)
	}

	return jsonData.Token, string(body), nil
}
