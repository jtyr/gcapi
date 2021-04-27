package grafana

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestGrafanaRestart(t *testing.T) {
	tests := []struct {
		body           string
		cfg            client.Config
		expectingError bool
		statusCode     int
	}{
		// Test good response
		{
			body:       "",
			statusCode: 200,
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
		// Test Not Found status code
		{
			expectingError: true,
			statusCode:     404,
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

		g := New()
		body, err := g.Restart()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Create: %s", i, err)
		}

		if !test.expectingError && test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}
	}
}
