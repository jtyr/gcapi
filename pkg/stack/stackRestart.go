package stack

import (
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Restart restarts a Stack and returns the raw API response.
func (s *stack) Restart() (string, error) {
	client, err := _client.New(ClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(s.endpoint + "/%s/restart", s.stackSlug)

	body, statusCode, err := client.Post(nil)
	if err != nil {
		if statusCode == 404 {
			return "", errors.New("Stack Slug not found")
		} else {
			return "", err
		}
	}

	return string(body), nil
}
