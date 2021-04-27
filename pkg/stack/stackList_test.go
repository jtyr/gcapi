package stack

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestStackList(t *testing.T) {
	tests := []struct {
		body             string
		cfg              client.Config
		expectedExitCode int
		expectingError   bool
		stackSlug        string
		statusCode       int
	}{
		// Test good response without Stack Slug
		{
			body: `{
					"items": [
						{
							"name": "testName",
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
						}
					]
				}`,
			expectedExitCode: consts.ExitOk,
			statusCode:       200,
		},
		// Test good response with Stack Slug
		{
			body: `{
					"items": [
						{
							"name": "testName",
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
						}
					]
				}`,
			expectedExitCode: consts.ExitOk,
			expectingError:   true,
			stackSlug:        "testSlug",
			statusCode:       200,
		},
		// Test bad JSON without Stack Slug
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       200,
		},
		// Test bad JSON with Stack Slug
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			stackSlug:        "testSlug",
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
		// Test Not Found without Stack Slug
		{
			expectedExitCode: consts.ExitError,
			expectingError:   true,
			statusCode:       404,
		},
		// Test Not Found with Stack Slug
		{
			expectedExitCode: consts.ExitNotFound,
			expectingError:   true,
			stackSlug:        "testSlug",
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

		st := New()
		if test.stackSlug != "" {
			st.SetStackSlug(test.stackSlug)
		}
		list, body, exitCode, err := st.List()
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
