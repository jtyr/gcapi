package grafana

import (
	"fmt"
	"strconv"
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
type grafana struct {
	// URL parameters
	stackSlug     string
	name          string
	role          string
	secondsToLive string

	// Relative path to the api-keys endpoint
	endpoint string
}

// ClientConfig holds the configuration for the HTTP Client
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *grafana {
	g := grafana{}

	g.endpoint = "instances/%s"

	return &g
}

// SetToken sets the authorization token used to communicate with the API.
func (g *grafana) SetToken(token string) {
	ClientConfig.Token = token
}

// SetStackSlug makes sure the Stack Slug is not an empty string.
func (g *grafana) SetStackSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Stack Slug value: "%s"`, value)
	}

	g.stackSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (g *grafana) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	g.name = value

	return nil
}

// SetRole makes sure the Role has correct value.
func (g *grafana) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		g.role = RoleAdmin
	case strings.ToLower(RoleEditor):
		g.role = RoleEditor
	case strings.ToLower(RoleViewer):
		g.role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}

// SetSecondsToLive makes sure the secondsToLive has correct value.
func (g *grafana) SetSecondsToLive(value uint64) error {
	if value > 0 {
		g.secondsToLive = strconv.FormatUint(value, 10)
	}

	return nil
}
