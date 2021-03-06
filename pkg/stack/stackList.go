package stack

import (
	"encoding/json"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
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
	TracesURL                string `json:"htInstanceUrl"`
	AlertManagerID           int    `json:"amInstanceId"`
	AlertManagerGeneratorURL string `json:"amInstanceGeneratorUrl"`
}

// listResp describes the structure of the JSON document returned by the API.
type listResp struct {
	Items []ListItem `json:"items"`
}

// List returns the list of API keys and raw API response.
func (s *Stack) List() (*[]ListItem, string, int, error) {
	client, err := _client.New(s.ClientConfig)
	if err != nil {
		return nil, "", consts.ExitError, fmt.Errorf("failed to get client: %s", err)
	}

	if s.StackSlug == "" {
		client.Endpoint = fmt.Sprintf("orgs/%s/"+s.Endpoint, s.OrgSlug)
	} else {
		client.Endpoint = fmt.Sprintf(s.Endpoint+"/%s", s.StackSlug)
	}

	body, statusCode, err := client.Get()
	if err != nil {
		if s.StackSlug == "" && statusCode == 404 {
			return nil, "", consts.ExitError, fmt.Errorf("Org Slug not found: %s", err)
		} else if s.StackSlug != "" && statusCode == 404 {
			return nil, "", consts.ExitNotFound, fmt.Errorf("Stack not found: %s", err)
		}

		return nil, "", consts.ExitError, err
	}

	var jsonData listResp

	if s.StackSlug != "" {
		jsonData.Items = append(jsonData.Items, ListItem{})

		if err := json.Unmarshal(body, &jsonData.Items[0]); err != nil {
			return nil, "", consts.ExitError, fmt.Errorf("cannot parse API response as JSON: %s", err)
		}
	} else {
		if err := json.Unmarshal(body, &jsonData); err != nil {
			return nil, "", consts.ExitError, fmt.Errorf("cannot parse API response as JSON: %s", err)
		}
	}

	return &jsonData.Items, string(body), consts.ExitOk, nil
}
