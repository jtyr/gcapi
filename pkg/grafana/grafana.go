package grafana

import (
	"fmt"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
)

const (
	// Allowed API key Role values expected by the API
	RoleAdmin  = "Admin"
	RoleEditor = "Editor"
	RoleViewer = "Viewer"
)

// grafana holds information about the Grafana.
type Grafana struct {
	// URL parameters
	StackSlug     string
	Name          string
	Role          string
	SecondsToLive string

	// Relative path to the api-keys endpoint
	Endpoint string
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *Grafana {
	g := Grafana{}

	g.Endpoint = "instances/%s"

	return &g
}

// SetToken sets the authorization token used to communicate with the API.
func (g *Grafana) SetToken(token string) {
	ClientConfig.Token = token
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
