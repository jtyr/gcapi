package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestClient(t *testing.T) {
	tests := []struct {
		body                 string
		cfg                  Config
		data                 interface{}
		env                  map[string]string
		expectingClientError bool
		expectingMethodError bool
		method               string
		statusCode           int
	}{
		// Test good GET response
		{
			body:       "OK",
			method:     "Get",
			statusCode: 200,
		},
		// Test bad URL
		{
			body: "OK",
			env: map[string]string{
				"GRAFANA_CLOUD_API_URL": "nothing",
			},
			expectingClientError: true,
			method:               "Get",
			statusCode:           200,
		},
		// Test bad Client
		{
			body: "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
			method:               "Get",
			statusCode:           400,
		},
		// Test good POST response
		{
			body:       "OK",
			method:     "Post",
			statusCode: 200,
		},
		// Test bad POST response
		{
			body: "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
			method:               "Post",
			statusCode:           400,
		},
		// Test good DELETE response
		{
			body:       "OK",
			method:     "Delete",
			statusCode: 200,
		},
		// Test bad DELETE response
		{
			body: "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
			method:               "Delete",
			statusCode:           400,
		},
	}

	for i, test := range tests {
		Client = tst.NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: test.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(test.body)),
			}
		})

		for k, v := range test.env {
			if err := os.Setenv(k, v); err != nil {
				t.Errorf("Test [%d] failed to set env var: %s", i, err)
			}
		}

		client, cerr := New(test.cfg)
		if test.expectingClientError && cerr != nil {
			for k := range test.env {
				if err := os.Unsetenv(k); err != nil {
					t.Errorf("Test [%d] failed to unset env var: %s", i, err)
				}
			}

			continue
		} else if cerr != nil {
			t.Errorf("Test [%d] failed to create New client: %s", i, cerr)
		}

		var body []byte
		var statusCode int
		var err error

		if test.method == "Get" {
			body, statusCode, err = client.Get()
		} else if test.method == "Post" {
			body, statusCode, err = client.Post(test.data)
		} else if test.method == "Delete" {
			body, statusCode, err = client.Delete()
		}

		if !test.expectingMethodError && err != nil {
			t.Errorf("Test [%d]: failed to run Get: %s", i, err)
		}

		if test.body != string(body) {
			t.Errorf("Test [%d]: expected body %s, got %s", i, test.body, string(body))
		}

		if test.statusCode != statusCode {
			t.Errorf("Test [%d]: expected status code %d, got %d", i, test.statusCode, statusCode)
		}
	}
}
