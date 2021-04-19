package stack

import (
	"fmt"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
)

// Stack holds information about the stack.
type Stack struct {
	// URL parameters
	OrgSlug   string
	StackSlug string
	Name      string

	// Relative path to the api-keys endpoint
	Endpoint string

	// HTTP client configuration
	ClientConfig client.Config
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.Config

// New returns new Stack.
func New() *Stack {
	s := Stack{}

	s.Endpoint = "instances"
	s.ClientConfig = ClientConfig

	return &s
}

// SetToken sets the authorization token used to communicate with the API.
func (s *Stack) SetToken(token string) {
	s.ClientConfig.Token = token
}

// SetOrgSlug makes sure the Org Slug is not an empty string.
func (s *Stack) SetOrgSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Org Slug value: "%s"`, value)
	}

	s.OrgSlug = value

	return nil
}

// SetStackSlug makes sure the Stack Slug is not an empty string.
func (s *Stack) SetStackSlug(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf(`invalid Stack Slug value: "%s"`, value)
	}

	s.StackSlug = value

	return nil
}

// SetName makes sure the Name is not an empty string.
func (s *Stack) SetName(value string) error {
	// TODO: Do further validation (only lowercase, no special chars?)
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("invalid Name value: %s", value)
	}

	s.Name = value

	return nil
}
