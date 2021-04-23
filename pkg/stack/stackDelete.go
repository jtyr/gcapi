package stack

import (
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
)

// Delete deletes the Stack.
func (s *Stack) Delete() (string, int, error) {
	client, err := _client.New(s.ClientConfig)
	if err != nil {
		return "", consts.ExitError, fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(s.Endpoint+"/%s", s.StackSlug)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", consts.ExitNotFound, fmt.Errorf("API key not found: %s", err)
		}

		return "", consts.ExitError, err
	}

	return string(body), consts.ExitOk, nil
}
