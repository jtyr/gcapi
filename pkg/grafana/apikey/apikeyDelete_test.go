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
		baseURL          string
		body             string
		cfg              client.Config
		deleteBody       string
		deleteClient     bool
		deleteStatusCode int
		expectedExitCode int
		expectingError   bool
		name             string
		statusCode       int
	}{
		// Test good response
		{
			baseURL:          "http://test",
			body:             `[{"id": 1, "name": "test", "role": "Viewer"}]`,
			deleteClient:     true,
			deleteStatusCode: 200,
			expectedExitCode: consts.ExitOk,
			name:             "test",
			statusCode:       200,
		},
		// Test bad KeyID
		{
			baseURL:          "http://test",
			body:             `[{"id": 1, "name": "badTest", "role": "Viewer"}]`,
			deleteClient:     true,
			deleteStatusCode: 200,
			expectedExitCode: consts.ExitNotFound,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
		},
		// Test bad List
		{
			baseURL:          "http://test",
			body:             `{}`,
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
		},
		// Test bad BaseURL
		{
			body:             `[{"id": 1, "name": "test", "role": "Viewer"}]`,
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
		},
		// Test bad Client
		{
			baseURL: "...",
			cfg: client.Config{
				Transport: &http.Transport{},
			},
			expectedExitCode: consts.ExitError,
			expectingError:   true,
		},
		// Test Not Found status code
		{
			baseURL:          "http://test",
			body:             `[{"id": 1, "name": "test", "role": "Viewer"}]`,
			deleteClient:     true,
			deleteStatusCode: 404,
			expectedExitCode: consts.ExitNotFound,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
		},
		// Test other status code
		{
			baseURL:          "http://test",
			body:             `[{"id": 1, "name": "test", "role": "Viewer"}]`,
			deleteClient:     true,
			deleteStatusCode: 429,
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			name:             "test",
			statusCode:       200,
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

		a := New()
		a.SetBaseURL(test.baseURL)
		a.SetName(test.name)

		if test.deleteClient {
			deleteClient = tst.NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: test.deleteStatusCode,
					Body:       ioutil.NopCloser(bytes.NewBufferString(test.deleteBody)),
				}
			})
		} else {
			deleteClient = &http.Client{}
		}

		body, exitCode, err := a.Delete()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Delete: %s", i, err)
		}

		if test.expectedExitCode != exitCode {
			t.Errorf("Test [%d]: expected Exit Code to be %d, got %d", i, test.expectedExitCode, exitCode)
		}

		if test.deleteBody != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.deleteBody, body)
		}
	}
}
