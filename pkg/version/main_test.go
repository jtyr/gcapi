package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version != "master" {
		t.Errorf("expected Version to be master, got %s", Version)
	}
}
