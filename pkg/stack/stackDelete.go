package stack

import (
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Delete deletes the Stack.
func (s *stack) Delete() (string, error) {
	client, err := _client.New(ClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(s.endpoint+"/%s", s.stackSlug)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", errors.New("API key with this name doesn't exist")
		} else {
			return "", err
		}
	}

	return string(body), nil
}
