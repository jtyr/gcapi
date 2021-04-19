package stack

import (
	"encoding/json"
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Create creates a new Stack and returns values of the newly created Stack and
// the raw API response.
func (s *Stack) Create() (*ListItem, string, error) {
	client, err := _client.New(s.ClientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = s.Endpoint

	var data []_client.Data
	data = append(data, _client.Data{Key: "name", Value: s.Name})
	data = append(data, _client.Data{Key: "slug", Value: s.StackSlug})

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return nil, "", errors.New("stack with this name already exists")
		}

		return nil, "", err
	}

	var jsonData ListItem
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, "", fmt.Errorf("cannot parse API response as JSON", err)
	}

	return &jsonData, string(body), nil
}
