package consts

import (
	"testing"
)

func TestConsts(t *testing.T) {
	if ExitOk != 0 {
		t.Errorf("expected ExitOk to be 0, got %d", ExitOk)
	}

	if ExitError != 1 {
		t.Errorf("expected ExitError to be 1, got %d", ExitError)
	}

	if ExitNotFound != 13 {
		t.Errorf("expected ExitNotFound to be 13, got %d", ExitNotFound)
	}
}
