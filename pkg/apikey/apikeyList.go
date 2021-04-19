package apikey

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jtyr/gcapi/pkg/client"
)

// ListItem described properties of individual List item returned by the API.
type ListItem struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// listResp describes the structure of the JSON document returned by the API.
type listResp struct {
	Items []ListItem `json:"items"`
}

// List returns the list of API keys and raw API response.
func (a *ApiKey) List() (*[]ListItem, string, error) {
	client, err := client.New(ClientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get client: %s", err)
	}

	if a.Name == "" {
		client.Endpoint = fmt.Sprintf(a.Endpoint, a.OrgSlug)
	} else {
		client.Endpoint = fmt.Sprintf(a.Endpoint+"/%s", a.OrgSlug, a.Name)
	}

	body, statusCode, err := client.Get()
	if err != nil {
		if statusCode == 404 {
			return nil, "", errors.New("Org Slug not found")
		} else {
			return nil, "", err
		}
	}

	var jsonData listResp
	if a.Name != "" {
		jsonData.Items = append(jsonData.Items, ListItem{})

		if err := json.Unmarshal(body, &jsonData.Items[0]); err != nil {
			return nil, "", fmt.Errorf("cannot parse API response as JSON", err)
		}
	} else {
		if err := json.Unmarshal(body, &jsonData); err != nil {
			return nil, "", fmt.Errorf("cannot parse API response as JSON", err)
		}
	}

	return &jsonData.Items, string(body), nil
}
