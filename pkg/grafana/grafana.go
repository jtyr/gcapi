package grafana

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
	_stack "github.com/jtyr/gcapi/pkg/stack"
)

// Grafana holds information about the Grafana.
type Grafana struct {
	// URL parameters
	OrgSlug   string
	StackSlug string
	Name      string

	// Relative path to the api-keys endpoint
	Endpoint        string
	GrafanaEndpoint string

	// API base address
	BaseURL string

	// HTTP client configuration
	ClientConfig client.Config

	// Grafana API token
	GrafanaToken string
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.Config

// New returns new Grafana.
func New() *Grafana {
	g := Grafana{}

	g.Endpoint = "instances/%s"
	g.ClientConfig = ClientConfig

	return &g
}

// SetToken sets the authorization token used to communicate with the API.
func (g *Grafana) SetToken(value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return errors.New("token has zero length")
	}

	g.ClientConfig.Token = value

	return nil
}

// SetGrafanaToken sets the authorization token used to communicate with the
// Grafana API.
func (g *Grafana) SetGrafanaToken(value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return errors.New("token has zero length")
	}

	g.GrafanaToken = value

	return nil
}

// SetOrgSlug makes sure the Org Slug is not an empty string.
func (g *Grafana) SetOrgSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Org Slug value: "%s"`, value)
	}

	g.OrgSlug = value

	return nil
}

// SetStackSlug makes sure the Stack Slug is not an empty string.
func (g *Grafana) SetStackSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Stack Slug value: "%s"`, value)
	}

	g.StackSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (g *Grafana) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	g.Name = value

	return nil
}

// SetBaseURL makes sure the BaseURL is not an empty string.
func (g *Grafana) SetBaseURL(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid BaseURL value: %s", value)
	}

	g.BaseURL = value

	return nil
}

// GetGrafanaAPIURL returns the URL used for Grafana
func (g *Grafana) GetGrafanaAPIURL() (string, error) {
	stack := _stack.New()

	if err := stack.SetOrgSlug(g.OrgSlug); err != nil {
		return "", err
	}

	if err := stack.SetStackSlug(g.StackSlug); err != nil {
		return "", err
	}

	stack.SetToken(g.ClientConfig.Token)

	list, _, _, err := stack.List()
	if err != nil {
		return "", fmt.Errorf("failed to get stack details: %s", err)
	}

	var grafanaURL string
	for _, item := range *list {
		grafanaURL = item.GrafanaURL + "/api"
	}

	return grafanaURL, nil
}
