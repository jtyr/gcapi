package apikey

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/grafana"
)

const (
	// Allowed API key Role values expected by the API
	RoleAdmin  = "Admin"
	RoleEditor = "Editor"
	RoleViewer = "Viewer"
)

// apiKey holds information about the Grafana API key.
type apiKey struct {
	// Inherit Grafana
	grafana.Grafana

	// URL parameters
	role          string
	secondsToLive string
}

// ClientConfig holds the configuration for the HTTP Client.
var ClientConfig client.ClientConfig

// New returns new ApiKey.
func New() *apiKey {
	ak := apiKey{}

	ak.Endpoint = "instances/%s/api/auth/keys"

	return &ak
}

// SetRole makes sure the Role has correct value.
func (ak *apiKey) SetRole(value string) error {
	switch strings.ToLower(value) {
	case strings.ToLower(RoleAdmin):
		ak.role = RoleAdmin
	case strings.ToLower(RoleEditor):
		ak.role = RoleEditor
	case strings.ToLower(RoleViewer):
		ak.role = RoleViewer
	default:
		return fmt.Errorf("invalid Role value: %s", value)
	}

	return nil
}

// SetSecondsToLive makes sure the secondsToLive has correct value.
func (ak *apiKey) SetSecondsToLive(value uint64) error {
	if value > 0 {
		ak.secondsToLive = strconv.FormatUint(value, 10)
	}

	return nil
}
