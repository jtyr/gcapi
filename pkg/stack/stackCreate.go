package stack

import (
	"encoding/json"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// createDocument describes the request JSON structure
type createDocument struct {
	Name      string `json:"name"`
	StackSlug string `json:"slug"`
}

// Create creates a new Stack and returns values of the newly created Stack and
// the raw API response.
func (s *Stack) Create() (*ListItem, string, error) {
	client, err := _client.New(s.ClientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = s.Endpoint

	data := createDocument{
		Name:      s.Name,
		StackSlug: s.StackSlug,
	}

	body, statusCode, err := client.Post(data)
	if err != nil {
		if statusCode == 409 {
			return nil, "", fmt.Errorf("stack with this name already exists: %s", err)
		}

		return nil, "", err
	}

	var jsonData ListItem
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, "", fmt.Errorf("cannot parse API response as JSON: %s", err)
	}

	return &jsonData, string(body), nil
}
