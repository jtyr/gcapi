package apikey

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestAPIKeyCreate(t *testing.T) {
	tests := []struct {
		body           string
		cfg            client.Config
		expectedToken  string
		expectingError bool
		statusCode     int
	}{
		// Test good response
		{
			body:          `{"token": "myToken"}`,
			expectedToken: "myToken",
			statusCode:    200,
		},
		// Test bad Client
		{
			expectingError: true,
			cfg: client.Config{
				Transport: &http.Transport{},
				BaseURL:   "...",
			},
			statusCode: 200,
		},
		// Test bad JSON
		{
			expectingError: true,
			statusCode:     200,
		},
		// Test Conflict status code
		{
			expectingError: true,
			statusCode:     409,
		},
		// Test other status code
		{
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

		ak := New()
		token, body, err := ak.Create()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Create: %s", i, err)
		}

		if test.expectedToken != token {
			t.Errorf("Test [%d]: expected Token to be %s, got %s", i, test.expectedToken, token)
		}

		if test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
