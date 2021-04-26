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
		statusCode           int
		body                 string
		method               string
		env                  map[string]string
		data                 interface{}
		expectingClientError bool
		expectingMethodError bool
		cfg                  Config
	}{
		{
			method:     "Get",
			statusCode: 200,
			body:       "OK",
		},
		{
			method:     "Get",
			statusCode: 200,
			body:       "OK",
			env: map[string]string{
				"GRAFANA_CLOUD_API_URL": "nothing",
			},
			expectingClientError: true,
		},
		{
			method:     "Get",
			statusCode: 400,
			body:       "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
		},
		{
			method:     "Post",
			statusCode: 200,
			body:       "OK",
		},
		{
			method:     "Post",
			statusCode: 400,
			body:       "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
		},
		{
			method:     "Delete",
			statusCode: 200,
			body:       "OK",
		},
		{
			method:     "Delete",
			statusCode: 400,
			body:       "",
			cfg: Config{
				Transport: &http.Transport{},
			},
			expectingMethodError: true,
		},
	}

	for i, test := range tests {
		Client = tst.NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: test.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(test.body)),
				Header:     make(http.Header),
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
