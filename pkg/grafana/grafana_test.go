package grafana

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jtyr/gcapi/pkg/client"
	tst "github.com/jtyr/gcapi/pkg/testing"
)

func TestGrafana(t *testing.T) {
	tests := []struct {
		baseURL                     string
		body                        string
		expectingBaseURLError       bool
		expectingGrafanaAPIURLError bool
		expectingGrafanaTokenError  bool
		expectingNameError          bool
		expectingOrgSlugError       bool
		expectingStackSlugError     bool
		expectingTokenError         bool
		grafanaAPIURL               string
		grafanaToken                string
		name                        string
		orgSlug                     string
		stackSlug                   string
		statusCode                  int
		token                       string
	}{
		// Test setting BaseURL, GrafanaAPIURL, GrafanaToken, Name, OrgSlug, StackSlug and Token
		{
			baseURL: "testBaseURL",
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
			grafanaAPIURL: "/api",
			grafanaToken:  "testGrafanaToken",
			name:          "testName",
			orgSlug:       "testOrgSlug",
			stackSlug:     "testStackSlug",
			statusCode:    200,
			token:         "testToken",
		},
		// Test GrafanaAPIURL with bad OrgSlug
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
			expectingBaseURLError:       true,
			expectingGrafanaAPIURLError: true,
			expectingGrafanaTokenError:  true,
			expectingNameError:          true,
			expectingOrgSlugError:       true,
			expectingStackSlugError:     true,
			expectingTokenError:         true,
			statusCode:                  200,
		},
		// Test GrafanaAPIURL with bad StackSlug
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
			expectingBaseURLError:       true,
			expectingGrafanaAPIURLError: true,
			expectingGrafanaTokenError:  true,
			expectingNameError:          true,
			expectingStackSlugError:     true,
			expectingTokenError:         true,
			orgSlug:                     "testOrgSlug",
			statusCode:                  200,
		},
		// Test GrafanaAPIURL with bad List
		{
			expectingBaseURLError:       true,
			expectingGrafanaAPIURLError: true,
			expectingGrafanaTokenError:  true,
			expectingNameError:          true,
			expectingTokenError:         true,
			orgSlug:                     "testOrgSlug",
			stackSlug:                   "testStackSlug",
			statusCode:                  200,
		},
		// Test incorrect BaseURL, GrafanaAPIURL, GrafanaToken, Name, OrgSlug, StackSlug and Token
		{
			expectingBaseURLError:       true,
			expectingGrafanaAPIURLError: true,
			expectingGrafanaTokenError:  true,
			expectingNameError:          true,
			expectingOrgSlugError:       true,
			expectingStackSlugError:     true,
			expectingTokenError:         true,
		},
	}

	for i, test := range tests {
		client.Client = tst.NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: test.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(test.body)),
			}
		})

		ClientConfig = client.Config{}

		g := New()
		g.OrgSlug = test.orgSlug
		g.StackSlug = test.stackSlug

		apiURL, err := g.GetGrafanaAPIURL()
		if err != nil {
			if !test.expectingGrafanaAPIURLError {
				t.Errorf("Test [%d]: failed to get Grafana API URL: %s", i, err)
			}
		} else if apiURL != test.grafanaAPIURL {
			t.Errorf("Test [%d]: expected Grafana API URL to be %s, got %s", i, test.grafanaAPIURL, apiURL)
		}

		if err := g.SetBaseURL(test.baseURL); err != nil {
			if !test.expectingBaseURLError {
				t.Errorf("Test [%d]: failed to set BaseURL: %s", i, err)
			}
		} else if g.BaseURL != test.baseURL {
			t.Errorf("Test [%d]: expected BaseURL to be %s, got %s", i, test.baseURL, g.BaseURL)
		}

		if err := g.SetGrafanaToken(test.grafanaToken); err != nil {
			if !test.expectingGrafanaTokenError {
				t.Errorf("Test [%d]: failed to set Grafana Token: %s", i, err)
			}
		} else if g.GrafanaToken != test.grafanaToken {
			t.Errorf("Test [%d]: expected Grafana Token to be %s, got %s", i, test.grafanaToken, g.GrafanaToken)
		}

		if err := g.SetName(test.name); err != nil {
			if !test.expectingNameError {
				t.Errorf("Test [%d]: failed to set Name: %s", i, err)
			}
		} else if g.Name != test.name {
			t.Errorf("Test [%d]: expected Name to be %s, got %s", i, test.name, g.Name)
		}

		if err := g.SetOrgSlug(test.orgSlug); err != nil {
			if !test.expectingOrgSlugError {
				t.Errorf("Test [%d]: failed to set Org Slug: %s", i, err)
			}
		} else if g.OrgSlug != test.orgSlug {
			t.Errorf("Test [%d]: expected Org Slug to be %s, got %s", i, test.orgSlug, g.OrgSlug)
		}

		if err := g.SetStackSlug(test.stackSlug); err != nil {
			if !test.expectingStackSlugError {
				t.Errorf("Test [%d]: failed to set Stack Slug: %s", i, err)
			}
		} else if g.StackSlug != test.stackSlug {
			t.Errorf("Test [%d]: expected Stack Slug to be %s, got %s", i, test.stackSlug, g.StackSlug)
		}

		if err := g.SetToken(test.token); err != nil {
			if !test.expectingTokenError {
				t.Errorf("Test [%d]: failed to set Token: %s", i, err)
			}
		} else if g.ClientConfig.Token != test.token {
			t.Errorf("Test [%d]: expected Token to be %s, got %s", i, test.token, g.ClientConfig.Token)
		}
	}
}
