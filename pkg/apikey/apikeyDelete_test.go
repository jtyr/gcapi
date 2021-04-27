package apikey

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestAPIKeyDelete(t *testing.T) {
	tests := []struct {
		body             string
		cfg              client.Config
		expectedExitCode int
		expectingError   bool
		statusCode       int
	}{
		// Test good response
		{
			expectedExitCode: consts.ExitOk,
			statusCode:       200,
		},
		// Test bad Client
		{
			cfg: client.Config{
				Transport: &http.Transport{},
				BaseURL:   "...",
			},
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       200,
		},
		// Test Not Found status code
		{
			expectedExitCode: consts.ExitNotFound,
			expectingError:   true,
			statusCode:       404,
		},
		// Test other status code
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       429,
		},
	}

	for i, test := range tests {
		client.Client = tst.NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: test.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(test.body)),
			}
		})

		if test.cfg.Transport != nil {
			ClientConfig = test.cfg
		} else {
			ClientConfig = client.Config{}
		}

		ak := New()
		body, exitCode, err := ak.Delete()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Delete: %s", i, err)
		}

		if test.expectedExitCode != exitCode {
			t.Errorf("Test [%d]: expected Exit Code to be %d, got %d", i, test.expectedExitCode, exitCode)
		}

		if test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
