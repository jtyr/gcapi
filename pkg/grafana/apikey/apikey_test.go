package apikey

import (
	"testing"
)

func TestApiKey(t *testing.T) {
	tests := []struct {
		expectedRole                string
		expectingRoleError          bool
		expectingSecondsToLiveError bool
		role                        string
		secondsToLive               uint64
	}{
		// Test setting "Admin" Role and SecondsToLive
		{
			expectedRole:  "Admin",
			role:          "admin",
			secondsToLive: 123,
		},
		// Test incorrect Role and SecondsToLive
		{
			expectingRoleError:          true,
			expectingSecondsToLiveError: true,
		},
		// Test "Editor" Role
		{
			expectedRole:                "Editor",
			expectingSecondsToLiveError: true,
			role:                        "editor",
		},
		// Test "Viewer" Role
		{
			expectedRole:                "Viewer",
			expectingSecondsToLiveError: true,
			role:                        "viewer",
		},
		// Test incorrect Role
		{
			expectingSecondsToLiveError: true,
			expectingRoleError:          true,
			role:                        "test",
		},
	}

	a := New()

	for i, test := range tests {
		if err := a.SetRole(test.role); err != nil {
			if !test.expectingRoleError {
				t.Errorf("Test [%d]: failed to set Role: %s", i, err)
			}
		} else if a.Role != test.expectedRole {
			t.Errorf("Test [%d]: expected Role to be %s, got %s", i, test.expectedRole, a.Role)
		}

		if err := a.SetSecondsToLive(test.secondsToLive); err != nil {
			if !test.expectingSecondsToLiveError {
				t.Errorf("Test [%d]: failed to set SecondsToLive: %s", i, err)
			}
		} else if a.SecondsToLive != test.secondsToLive {
			t.Errorf("Test [%d]: expected SecondsToLive to be %d, got %d", i, test.secondsToLive, a.SecondsToLive)
		}
	}
}
