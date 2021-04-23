package apikey

import (
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
	"github.com/jtyr/gcapi/pkg/consts"
)

// Delete deletes the API key.
func (a *APIKey) Delete() (string, int, error) {
	client, err := _client.New(a.ClientConfig)
	if err != nil {
		return "", consts.ExitError, fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(a.Endpoint+"/%s", a.OrgSlug, a.Name)

	body, statusCode, err := client.Delete()
	if err != nil {
		if statusCode == 404 {
			return "", consts.ExitNotFound, fmt.Errorf("API key not found: %s", err)
		}

		return "", consts.ExitError, err
	}

	return string(body), consts.ExitOk, nil
}
