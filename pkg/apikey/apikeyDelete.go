package apikey

import (
	"errors"
	"fmt"

	_client "github.com/jtyr/gcapi/pkg/client"
)

// Delete deletes the API key.
func (a *ApiKey) Delete() (string, error) {
	client, err := _client.New(a.ClientConfig)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %s", err)
	}

	client.Endpoint = fmt.Sprintf(a.Endpoint+"/%s", a.OrgSlug, a.Name)

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
