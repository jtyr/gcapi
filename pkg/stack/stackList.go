package stack

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jtyr/gcapi/pkg/client"
)

// ListItem described properties of individual List item returned by the API.
type ListItem struct {
	Name                     string `json:"name"`
	Slug                     string `json:"slug"`
	GrafanaURL               string `json:"url"`
	PrometheusID             int    `json:"hmInstancePromId"`
	PrometheusURL            string `json:"hmInstancePromUrl"`
	GraphiteID               int    `json:"hmInstanceGraphiteId"`
	GraphiteURL              string `json:"hmInstanceGraphiteUrl"`
	LogsID                   int    `json:"hlInstanceId"`
	LogsURL                  string `json:"hlInstanceUrl"`
	TracesID                 int    `json:"htInstanceId"`
	AlertManagerID           int    `json:"amInstanceId"`
	AlertManagerGeneratorURL string `json:"amInstanceGeneratorUrl"`
}

// listResp describes the structure of the JSON document returned by the API.
type listResp struct {
	Items []ListItem `json:"items"`
}

// List returns the list of API keys and raw API response.
func (s *Stack) List() (*[]ListItem, string, error) {
	client, err := client.New(s.ClientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get client: %s", err)
	}

	if s.StackSlug == "" {
		client.Endpoint = fmt.Sprintf("orgs/%s/"+s.Endpoint, s.OrgSlug)
	} else {
		client.Endpoint = fmt.Sprintf(s.Endpoint+"/%s", s.StackSlug)
	}

	body, statusCode, err := client.Get()
	if err != nil {
		if statusCode == 404 {
			return nil, "", errors.New("Stack Slug not found")
		} else {
			return nil, "", err
		}
	}

	var jsonData listResp

	if s.StackSlug != "" {
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
