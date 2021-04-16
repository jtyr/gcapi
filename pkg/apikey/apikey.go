package apikey

import (
	"fmt"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
)

const (
	// Allowed values expected by the API
	RoleViewer           = "Viewer"
	RoleEditor           = "Editor"
	RoleMetricsPublisher = "MetricsPublisher"
	RoleAdmin            = "Admin"
)

// apiKey holds information about the API key.
type apiKey struct {
	// URL parameters
	orgSlug string
	name    string
	role    string

	// Relative path to the api-keys endpoint
	endpoint string
}

// ClientConfig holds the configuration for the HTTP Client
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *apiKey {
	a := apiKey{}

	a.endpoint = "orgs/%s/api-keys"

	return &a
}

// SetToken sets the authorization token used to communicate with the API.
func (a *apiKey) SetToken(token string) {
	ClientConfig.Token = token
}

// SetOrgSlug makes sure the Org Slug is not an empty string.
func (a *apiKey) SetOrgSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Org Slug value: "%s"`, value)
	}

	a.orgSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (a *apiKey) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	a.name = value

	return nil
}

// SetRole makes sure the role has correct value.
func (a *apiKey) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		a.role = RoleAdmin
	case strings.ToLower(RoleEditor):
		a.role = RoleEditor
	case strings.ToLower(RoleMetricsPublisher):
		a.role = RoleMetricsPublisher
	case strings.ToLower(RoleViewer):
		a.role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}
