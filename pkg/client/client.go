package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_url "net/url"
	"os"
)

// Config allows to override default HTTP client configuration.
type Config struct {
	// BaseURL is API base address
	BaseURL string

	// Token is authorization token
	Token string

	// Transport allows to override the default HTTP transport of the client
	Transport *http.Transport
}

// GrafanaCloudClient holds the HTTP client configuration.
type GrafanaCloudClient struct {
	// Client refers to the HTTP client
	Client *http.Client
	// Endpoint os relative URL to the API endpoint
	Endpoint string

	// token ia authorization token
	token string

	// baseURL is base URL used to construct the full endpoint URL
	baseURL *_url.URL
}

// Data defines the POST data key=value pair.
type Data struct {
	Key   string
	Value string
}

// Client allows to mock the http.client from tests.
var Client = &http.Client{}

// New creates new client.
func New(cfg Config) (*GrafanaCloudClient, error) {
	c := GrafanaCloudClient{}
	c.token = cfg.Token
	c.Client = Client

	if cfg.Transport != nil {
		c.Client.Transport = cfg.Transport
	}

	envURL := os.Getenv("GRAFANA_CLOUD_API_URL")
	if envURL != "" {
		cfg.BaseURL = envURL
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://grafana.com/api"
	}

	var err error
	c.baseURL, err = _url.ParseRequestURI(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse BaseURL: %s", err)
	}

	return &c, nil
}

// Get sends GET request.
func (c *GrafanaCloudClient) Get() ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, c.Endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.do(req)
}

// Post sends POST request.
func (c *GrafanaCloudClient) Post(data interface{}) ([]byte, int, error) {
	jsonDoc, err := json.Marshal(data)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to convert struct to JSON: %s", err)
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, c.Endpoint)

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonDoc))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}

// Delete sends DELETE request.
func (c *GrafanaCloudClient) Delete() ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, c.Endpoint)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.do(req)
}

func (c *GrafanaCloudClient) do(req *http.Request) ([]byte, int, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, fmt.Errorf("bad status code received: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error reading response: %s", err)
	}

	return body, resp.StatusCode, nil
}
