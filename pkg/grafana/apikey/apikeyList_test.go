package apikey

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestAPIKeyList(t *testing.T) {
	tests := []struct {
		baseURL        string
		body           string
		cfg            client.Config
		expectingError bool
		statusCode     int
	}{
		// Test good response
		{
			baseURL:    "http://test",
			body:       `[{"id": 1, "name": "test", "role": "Viewer"}]`,
			statusCode: 200,
		},
		// Test bad BaseURL
		{
			expectingError: true,
		},
		// Test bad JSON without Name
		{
			baseURL:        "http://test",
			body:           "",
			expectingError: true,
			statusCode:     200,
		},
		// Test bad Client
		{
			baseURL: "...",
			cfg: client.Config{
				Transport: &http.Transport{},
			},
			expectingError: true,
		},
		// Test Not Found status code
		{
			baseURL:        "http://test",
			expectingError: true,
			statusCode:     404,
		},
		// Test other status code
		{
			baseURL:        "http://test",
			expectingError: true,
			statusCode:     429,
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

		list, body, ec, err := a.List()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call List (exit code=%d): %s", i, ec, err)
		}

		if list != nil && len(*list) == 0 {
			t.Errorf("Test [%d]: expected list to have non-zero length", i)
		}

		if test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
