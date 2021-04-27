package stack

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestStackCreate(t *testing.T) {
	tests := []struct {
		body           string
		cfg            client.Config
		expectingError bool
		statusCode     int
	}{
		// Test good response
		{
			body: `{"name": "testName",
				"slug": "testSlug",
				"url": "testUrl",
				"hmInstancePromId": 123,
				"hmInstancePromUrl": "metricsUrl",
				"hmInstanceGraphiteId": 456,
				"hmInstanceGraphiteUrl": "graphiteUrl",
				"hlInstanceId": 789,
				"hlInstanceUrl": "logsUrl",
				"htInstanceId": 987,
				"amInstanceId": 654,
				"amInstanceGeneratorUrl": "alertsUrl"
				}`,
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

		st := New()
		item, body, err := st.Create()
		if !test.expectingError && err != nil {
			t.Errorf("Test [%d]: failed to call Create: %s", i, err)
		}

		if !test.expectingError && test.body != body {
			t.Errorf("Test [%d]: expected Body to be %s, got %s", i, test.body, body)
		}

		if !test.expectingError && item == nil {
			t.Errorf("Test [%d]: List Item is not set", i)
		}
	}
}
