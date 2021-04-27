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

func TestAPIKeyList(t *testing.T) {
	tests := []struct {
		body             string
		cfg              client.Config
		expectedExitCode int
		expectingError   bool
		name             string
		statusCode       int
	}{
		// Test good response without Name
		{
			body:             `{"items": [{"name": "test", "role": "viewer"}]}`,
			expectedExitCode: consts.ExitOk,
			statusCode:       200,
		},
		// Test good response with Name
		{
			body:             `{"items": [{"name": "test", "role": "viewer"}]}`,
			expectedExitCode: consts.ExitOk,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
		},
		// Test bad JSON without Name
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       200,
		},
		// Test bad JSON with Name
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			name:             "test",
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
		// Test Not Found without Name
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       404,
		},
		// Test Not Found with Name
		{
			expectedExitCode: consts.ExitNotFound,
			expectingError:   true,
			name:             "test",
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
		if test.name != "" {
			ak.SetName(test.name)
		}
		list, body, exitCode, err := ak.List()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call List: %s", i, err)
		}

		if list != nil && len(*list) == 0 {
			t.Errorf("Test [%d]: expected list to have non-zero length", i)
		}

		if test.expectedExitCode != exitCode {
			t.Errorf("Test [%d]: expected Exit Code to be %d, got %d", i, test.expectedExitCode, exitCode)
		}

		if test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
