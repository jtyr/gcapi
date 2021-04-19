package stack

import (
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Delete deletes the Stack.
func (s *Stack) Delete() (string, error) {
	client, err := _client.New(s.ClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(s.Endpoint+"/%s", s.StackSlug)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", errors.New("API key not found")
		} else {
			return "", err
		}
	}

	return string(body), nil
}
