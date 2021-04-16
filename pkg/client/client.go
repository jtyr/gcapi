package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_url "net/url"
	"strings"
)

// ClientConfig allows to override default HTTP client configuration.
type ClientConfig struct {
	// API base address
	BaseURL string

	// Authorization token
	Token string

	// Allows to override the default HTTP transport of the client
	Transport http.Transport
}

// GrafanaCloudClient holds the HTTP client configuration.
type GrafanaCloudClient struct {
	// HTTP client
	Client *http.Client
	// Relative URL to the API endpoint
	Endpoint string

	// Authorization token
	token string

	// Base URL used to construct the full endpoint URL
	baseURL *_url.URL
}

// Data defines the POST data key=value pair.
type Data struct {
	Key   string
	Value string
}

// New creates new client.
func New(cfg ClientConfig) (*GrafanaCloudClient, error) {
	c := GrafanaCloudClient{}

	c.token = cfg.Token

	c.Client = &http.Client{
		Transport: &cfg.Transport,
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

// Post sends POST request.
func (c *GrafanaCloudClient) Post(data []Data) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, c.Endpoint)

	req, err := http.NewRequest("POST", url, c.getDataReader(data))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

// Delete sends DELETE request.
func (c *GrafanaCloudClient) Delete() ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, c.Endpoint)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

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

// getDataReader returns Data as Reader ready to used in the http.NewRequest.
func (c GrafanaCloudClient) getDataReader(data []Data) *strings.Reader {
	dataString := ""
	dataLen := len(data)

	for i, kv := range data {
		dataString += kv.Key + "=" + _url.QueryEscape(kv.Value)

		if i < dataLen-1 {
			dataString += "&"
		}
	}

	return strings.NewReader(dataString)
}
