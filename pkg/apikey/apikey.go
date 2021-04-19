package apikey

import (
	"fmt"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
)

const (
	// Allowed API key Role values expected by the API
	RoleAdmin            = "Admin"
	RoleEditor           = "Editor"
	RoleMetricsPublisher = "MetricsPublisher"
	RoleViewer           = "Viewer"
)

// ApiKey holds information about the API key.
type ApiKey struct {
	// URL parameters
	OrgSlug string
	Name    string
	Role    string

	// Relative path to the api-keys endpoint
	Endpoint string

	// HTTP client configuration
	ClientConfig client.ClientConfig
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *ApiKey {
	a := ApiKey{}

	a.Endpoint = "orgs/%s/api-keys"
	a.ClientConfig = ClientConfig

	return &a
}

// SetToken sets the authorization token used to communicate with the API.
func (a *ApiKey) SetToken(token string) {
	a.ClientConfig.Token = token
}

// SetOrgSlug makes sure the Org Slug is not an empty string.
func (a *ApiKey) SetOrgSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Org Slug value: "%s"`, value)
	}

	a.OrgSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (a *ApiKey) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	a.Name = value

	return nil
}

// SetRole makes sure the role has correct value.
func (a *ApiKey) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		a.Role = RoleAdmin
	case strings.ToLower(RoleEditor):
		a.Role = RoleEditor
	case strings.ToLower(RoleMetricsPublisher):
		a.Role = RoleMetricsPublisher
	case strings.ToLower(RoleViewer):
		a.Role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}
