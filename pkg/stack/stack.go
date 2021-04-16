package stack

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

// stack holds information about the stack.
type stack struct {
	// URL parameters
	orgSlug   string
	stackSlug string
	name      string
	role      string

	// Relative path to the api-keys endpoint
	endpoint string
}

// ClientConfig holds the configuration for the HTTP Client
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *stack {
	s := stack{}

	s.endpoint = "instances"

	return &s
}

// SetToken sets the authorization token used to communicate with the API.
func (s *stack) SetToken(token string) {
	ClientConfig.Token = token
}

// SetOrgSlug makes sure the Org Slug is not an empty string.
func (s *stack) SetOrgSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Org Slug value: "%s"`, value)
	}

	s.orgSlug = value

	return nil
}

// SetStackSlug makes sure the Stack Slug is not an empty string.
func (s *stack) SetStackSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Stack Slug value: "%s"`, value)
	}

	s.stackSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (s *stack) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	s.name = value

	return nil
}

// SetRole makes sure the role has correct value.
func (s *stack) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		s.role = RoleAdmin
	case strings.ToLower(RoleEditor):
		s.role = RoleEditor
	case strings.ToLower(RoleMetricsPublisher):
		s.role = RoleMetricsPublisher
	case strings.ToLower(RoleViewer):
		s.role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}
