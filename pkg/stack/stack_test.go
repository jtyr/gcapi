package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		expectingNameError      bool
		expectingOrgSlugError   bool
		expectingStackSlugError bool
		expectingTokenError     bool
		name                    string
		orgSlug                 string
		stackSlug               string
		token                   string
	}{
		// Test setting Name, OrgSlug, StackSlug and Token
		{
			name:      "testName",
			orgSlug:   "testOrgSlug",
			stackSlug: "test",
			token:     "test token",
		},
		// Test incorrect Name, OrgSlug, StackSlug and Token
		{
			expectingNameError:      true,
			expectingOrgSlugError:   true,
			expectingStackSlugError: true,
			expectingTokenError:     true,
		},
	}

	st := New()

	for i, test := range tests {
		if err := st.SetName(test.name); err != nil {
			if !test.expectingNameError {
				t.Errorf("Test [%d]: failed to set Name: %s", i, err)
			}
		} else if st.Name != test.name {
			t.Errorf("Test [%d]: expected Name to be %s, got %s", i, test.name, st.Name)
		}

		if err := st.SetOrgSlug(test.orgSlug); err != nil {
			if !test.expectingOrgSlugError {
				t.Errorf("Test [%d]: failed to set Org Slug: %s", i, err)
			}
		} else if st.OrgSlug != test.orgSlug {
			t.Errorf("Test [%d]: expected Org Slug to be %s, got %s", i, test.orgSlug, st.OrgSlug)
		}

		if err := st.SetStackSlug(test.stackSlug); err != nil {
			if !test.expectingStackSlugError {
				t.Errorf("Test [%d]: failed to set Stack Slug: %s", i, err)
			}
		} else if st.StackSlug != test.stackSlug {
			t.Errorf("Test [%d]: expected Stack Slug to be %s, got %s", i, test.stackSlug, st.StackSlug)
		}

		if err := st.SetToken(test.token); err != nil {
			if !test.expectingTokenError {
				t.Errorf("Test [%d]: failed to set Token: %s", i, err)
			}
		} else if st.ClientConfig.Token != test.token {
			t.Errorf("Test [%d]: expected Token to be %s, got %s", i, test.token, st.ClientConfig.Token)
		}
	}
}
