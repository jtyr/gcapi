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
		baseURL        string
		body           string
		cfg            client.Config
		expectedKey    string
		expectingError bool
		orgSlug        string
		secondsToLive  uint64
		stackSlug      string
		statusCode     int
	}{
		// Test good response without BaseURL
		{
			body:        `{"key": "myKey"}`,
			expectedKey: "myKey",
			orgSlug:     "testOrgSlug",
			stackSlug:   "testStackSlug",
			statusCode:  200,
		},
		// Test good response with BaseURL
		{
			baseURL:     "http://test",
			body:        `{"key": "myKey"}`,
			expectedKey: "myKey",
			orgSlug:     "testOrgSlug",
			stackSlug:   "testStackSlug",
			statusCode:  200,
		},
		// Test good response with SecondsToLive
		{
			baseURL:       "http://test",
			body:          `{"key": "myKey"}`,
			expectedKey:   "myKey",
			orgSlug:       "testOrgSlug",
			stackSlug:     "testStackSlug",
			secondsToLive: 123,
			statusCode:    200,
		},
		// Test bad Client
		{
			baseURL:        "...",
			expectingError: true,
			cfg: client.Config{
				Transport: &http.Transport{},
			},
			statusCode: 200,
		},
		// Test empty BaseURL
		{
			expectingError: true,
			statusCode:     200,
		},
		// Test empty StackSlug
		{
			baseURL:        "http://test",
			expectingError: true,
			statusCode:     200,
		},
		// Test Conflict status code
		{
			baseURL:        "http://test",
			expectingError: true,
			statusCode:     409,
		},
		// Test other status code
		{
			baseURL:        "http://localhost",
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
		a.SetOrgSlug(test.orgSlug)
		a.SetStackSlug(test.stackSlug)
		a.SetBaseURL(test.baseURL)
		a.SetSecondsToLive(test.secondsToLive)

		key, body, err := a.Create()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Create: %s", i, err)
		}

		if test.expectedKey != key {
			t.Errorf("Test [%d]: expected Key to be %s, got %s", i, test.expectedKey, key)
		}

		if test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
