[![License](https://img.shields.io/github/license/jtyr/gcapi)](LICENSE)
[![Actions status](https://github.com/jtyr/gcapi/actions/workflows/go.yaml/badge.svg)](https://github.com/jtyr/gcapi/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtyr/gcapi)](https://goreportcard.com/report/github.com/jtyr/gcapi)


Grafana Cloud API
=================

Golang package that allows an easy access to the [Grafana Cloud
API](https://grafana.com/docs/grafana-cloud/api/) (and [Grafana
API](https://grafana.com/docs/grafana/latest/http_api/)).


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
  apikey      Manage Grafana Cloud API keys
  grafana     Manage Grafana instance
  help        Help about any command
  stack       Manage Grafana Cloud stacks
  version     Show version

Flags:
  -t, --cloud-api-token string        Grafana Cloud API token
  -f, --cloud-api-token-file string   path to a file containing the Grafana Cloud API token
  -h, --help                          help for gcapi-cli
      --timestamps                    enable Log timestamps
  -v, --version                       version for gcapi-cli

Use "gcapi-cli [command] --help" for more information about a command.
```

Examples:

```shell
# Export the authorization token for the Grafana Cloud API
export GRAFANA_CLOUD_API_TOKEN='abcdefghijklmnopqrstuvwxyz0123456789'
# Export Grafana API URL
# (needed only for some of the "gcapi-cli grafana" subcommands)
export GRAFANA_API_URL='http://grafana.domain.com/api'
# Export the authorization token for the Grafana API
# (use "gcapi-cli grafana apikey create" to create one)
# (needed only for some of the "gcapi-cli grafana" subcommands)
export GRAFANA_API_TOKEN='9876543210zyxwvutsrqponmlkjihgfedcba'

#####
### API key
#####

# Create a new API key
gcapi-cli apikey create myorgslug myname
# List all API keys
gcapi-cli apikey list myorgslug
# List a specific API key
gcapi-cli apikey list myorgslug myname
# Delete an API key
gcapi-cli apikey delete myorgslug myname

#####
### Stack
#####

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

#####
### Grafana
#####

# Restart Grafana
gcapi-cli grafana restart mystackslug

## Grafana API key (using Grafana Cloud API and Grafana API)
# Create Grafana API key
gcapi-cli grafana apikey create mystackslug myname Viewer
# List Grafana API keys
gcapi-cli grafana apikey list myorgslug mystackslug
# List a specific Grafana API key
gcapi-cli grafana apikey list myorgslug mystackslug myname
# Delete a Grafana API key
gcapi-cli grafana apikey delete myorgslug mystackslug myname

## Grafana API key (using only Grafana API)
# Create Grafana API key
gcapi-cli grafana apikey create myname Viewer
# List Grafana API keys
gcapi-cli grafana apikey list
# List a specific Grafana API key
gcapi-cli grafana apikey list myname
# Delete a Grafana API key
gcapi-cli grafana apikey delete myname

#####
 ### TODO
  #

## Grafana Datasource
# Create a Grafana Datasource
gcapi-cli grafana datasource create myorgslug mystackslug /path/to/my/datasource.json
# List Grafana Datasources
gcapi-cli grafana datasource list mystackslug
# Delete a Grafana Datasource
gcapi-cli grafana datasource delete mystackslug myid

## Grafana Dashboard
# Create a Grafana Dashboard
gcapi-cli grafana datasource create myorgslug mystackslug /path/to/my/dashboard.json
# List Grafana Dashboards
gcapi-cli grafana datasource list myorgslug mystackslug
# Delete a Grafana Dashboard
gcapi-cli grafana datasource delete myorgslug mystackslug myid
```

Environment variables:

- `GRAFANA_API_TOKEN` - Authorization token used for communication with the
     [Grafana API](https://grafana.com/docs/grafana/latest/http_api/). This
     variable overrides `--grafana-api-token` and `--grafana-api-token-file` if
     they are specified. Needed only for some of the `gcapi-cli grafana`
     subcommands that are using the [Grafana
     API](https://grafana.com/docs/grafana/latest/http_api/).
- `GRAFANA_API_URL` - URL to the [Grafana
     API](https://grafana.com/docs/grafana/latest/http_api/) (e.g.
     `http://grafana.domain.com/api`). This variable overrides
     `--grafana-api-url` if specified. Needed only for some of the `gcapi-cli
     grafana` subcommands that are using the [Grafana
     API](https://grafana.com/docs/grafana/latest/http_api/).
- `GRAFANA_CLOUD_API_TOKEN` - Authorization token used for communication with
     the [Grafana Cloud API](https://grafana.com/docs/grafana-cloud/api/). This
     variable overrides `--cloud-api-token` and `--cloud-api-token-file` if
     specified.
- `GRAFANA_CLOUD_API_URL` - URL to the [Grafana Cloud
     API](https://grafana.com/docs/grafana-cloud/api/) (default
     `https://grafana.com/api`).


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


TODO
----

- Testing (help wanted)


Author
------

Jiri Tyr


License
-------

Apache 2.0
