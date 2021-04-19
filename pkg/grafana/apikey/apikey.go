package apikey

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/grafana"
)

// Allowed API key Role values expected by the API
const (
	RoleAdmin  = "Admin"
	RoleEditor = "Editor"
	RoleViewer = "Viewer"
)

// APIKey holds information about the Grafana API key.
type APIKey struct {
	// Inherit Grafana
	grafana.Grafana

	// URL parameters
	Role          string
	SecondsToLive string
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.Config

// New returns new APIKey.
func New() *APIKey {
	a := APIKey{}

	// Set API endpoints
	a.Endpoint = "instances/%s/api/auth/keys"
	a.GrafanaEndpoint = "auth/keys"

	// Set HTTP client config
	a.ClientConfig = ClientConfig

	return &a
}

// SetRole makes sure the Role has correct value.
func (a *APIKey) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		a.Role = RoleAdmin
	case strings.ToLower(RoleEditor):
		a.Role = RoleEditor
	case strings.ToLower(RoleViewer):
		a.Role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}

// SetSecondsToLive makes sure the secondsToLive has correct value.
func (a *APIKey) SetSecondsToLive(value uint64) error {
	if value > 0 {
		a.SecondsToLive = strconv.FormatUint(value, 10)
	}

	return nil
}
