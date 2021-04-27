package apikey

import (
	"testing"
)

func TestApiKey(t *testing.T) {
	tests := []struct {
		expectedRole          string
		expectingNameError    bool
		expectingOrgSlugError bool
		expectingRoleError    bool
		expectingTokenError   bool
		name                  string
		orgSlug               string
		role                  string
		token                 string
	}{
		// Test setting Name, OrgSlug, Token and "Admin" Role
		{
			expectedRole: "Admin",
			name:         "testName",
			orgSlug:      "testOrgSlug",
			role:         "admin",
			token:        "testToken",
		},
		// Test incorrect Name, OrgSlug, Token and Role
		{
			expectingNameError:    true,
			expectingOrgSlugError: true,
			expectingRoleError:    true,
			expectingTokenError:   true,
		},
		// Test "Editor" Role
		{
			expectedRole:          "Editor",
			expectingNameError:    true,
			expectingOrgSlugError: true,
			expectingTokenError:   true,
			role:                  "editor",
		},
		// Test "MetricsPublisher" Role
		{
			expectedRole:          "MetricsPublisher",
			expectingNameError:    true,
			expectingOrgSlugError: true,
			expectingTokenError:   true,
			role:                  "metricspublisher",
		},
		// Test "Viewer" Role
		{
			expectedRole:          "Viewer",
			expectingNameError:    true,
			expectingOrgSlugError: true,
			expectingTokenError:   true,
			role:                  "viewer",
		},
		// Test incorrect Role
		{
			expectingNameError:    true,
			expectingOrgSlugError: true,
			expectingRoleError:    true,
			expectingTokenError:   true,
			role:                  "test",
		},
	}

	ak := New()

	for i, test := range tests {
		if err := ak.SetName(test.name); err != nil {
			if !test.expectingNameError {
				t.Errorf("Test [%d]: failed to set Name: %s", i, err)
			}
		} else if ak.Name != test.name {
			t.Errorf("Test [%d]: expected Name to be %s, got %s", i, test.name, ak.Name)
		}

		if err := ak.SetOrgSlug(test.orgSlug); err != nil {
			if !test.expectingOrgSlugError {
				t.Errorf("Test [%d]: failed to set Org Slug: %s", i, err)
			}
		} else if ak.OrgSlug != test.orgSlug {
			t.Errorf("Test [%d]: expected Org Slug to be %s, got %s", i, test.orgSlug, ak.OrgSlug)
		}

		if err := ak.SetRole(test.role); err != nil {
			if !test.expectingRoleError {
				t.Errorf("Test [%d]: failed to set Role: %s", i, err)
			}
		} else if ak.Role != test.expectedRole {
			t.Errorf("Test [%d]: expected Role to be %s, got %s", i, test.expectedRole, ak.Role)
		}

		if err := ak.SetToken(test.token); err != nil {
			if !test.expectingTokenError {
				t.Errorf("Test [%d]: failed to set Token: %s", i, err)
			}
		} else if ak.ClientConfig.Token != test.token {
			t.Errorf("Test [%d]: expected Token to be %s, got %s", i, test.token, ak.ClientConfig.Token)
		}
	}
}
