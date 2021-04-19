Grafana Cloud API
=================

Golang package that allows an easy access to the Grafana Cloud API.


Compilation
-----------

```shell
go mod vendor
go build -o gcapi-cli ./main.go
```


Usage
-----

### Command line tool

```
$ gcapi-cli --help
gcapi-cli is a tool that allows an easy access to the Grafana Cloud API.

Usage:
  gcapi-cli [flags]
  gcapi-cli [command]

Available Commands:
  apikey      Manage API keys
  help        Help about any command
  stack       Manage Stacks
  version     Show version

Flags:
  -t, --api-token string        Grafana Cloud API token
  -f, --api-token-file string   path to a file containing the Grafana Cloud API token
  -h, --help                    help for gcapi-cli
      --timestamps              enable Log timestamps
  -v, --version                 version for gcapi-cli

Use "gcapi-cli [command] --help" for more information about a command.
```

Examples:

```shell
# Export the authorization token
export GRAFANA_CLOUD_API_TOKEN='abcdefghijklmnopqrstuvwxyz0123456789'

### API key
# Create a new API key
gcapi-cli apikey create myorgslug myname
# List all API keys
gcapi-cli apikey list myorgslug
# List a specific API key
gcapi-cli apikey list myorgslug myname
# Delete an API key
gcapi-cli apikey delete myorgslug myname

### Stack
# Create a new Stack if the Stack Slug is the same as the Name
gcapi-cli stack create mystackslug
# Create a new Stack if the Stack Slug is different from the Name
gcapi-cli stack create mystackslug myname
# List all Stacks
gcapi-cli stack list myorgslug
# List a specific Stack
gcapi-cli stack list myorgslug myname
# Delete an API key
gcapi-cli stack delete mystackslug

### Grafana
# Restart Grafana
gcapi-cli grafana restart mystackslug
# Create Grafana API key
gcapi-cli grafana apikey create mystackslug myname Viewer

### TODO
# Create a Grafana Datasource
gcapi-cli grafana datasource create mystackslug /path/to/my/datasource.json
# List Grafana Datasources
gcapi-cli grafana datasource list mystackslug
# Delete a Grafana Datasource
gcapi-cli grafana datasource delete mystackslug mydsid
# List Grafana API keys
gcapi-cli grafana apikey list mystackslug
# List a specific Grafana API key
gcapi-cli grafana apikey list mystackslug myname
# Delete a Grafana API key
gcapi-cli grafana apikey delete mystackslug myname
```

Environment variables:

`GRAFANA_CLOUD_API_URL` - URL to the Grafana Cloud API (default `https://grafana.com/api`).
`GRAFANA_CLOUD_API_TOKEN` - Authorization token used for communication with the Grafana Cloud API.


### Go package

```go
import (
	"fmt"
	"os"

	"github.com/jtyr/gcapi/pkg/apikey"
)

func main() {
	// Get new API key instance
	ak = apikey.New()

	// Set authorization token
	ak.SetToken("abcdefghijklmnopqrstuvwxyz0123456789")

	// Set the API key parameters
	ak.SetOrgSlug("myorgslug")
	ak.SetName("myname")
	ak.SetRole(apikey.RoleViewer)

	// Create new API key
	key, err := ak.Create()
	if err != nil {
		fmt.Printf("failed to create API key: %s", err)
		os.Exit(1)
	}

	// Print the created API key
	fmt.Printf("New API key is: %s\n", key)
}
```


Author
------

Jiri Tyr


License
-------

Apache 2.0
