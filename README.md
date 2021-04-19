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
# Restart a Stack
gcapi-cli stack restart mystackslug

### TODO
# Create Stack API key
gcapi-cli stack apikey create mystackslug myname Viewer
# Create Stack API key with limitted life
gcapi-cli stack apikey create mystackslug myname Viewer --secondsToLive 3600
# List Stack datasources
gcapi-cli stack datasource list mystackslug
# Delete a Stack datasource
gcapi-cli stack datasource delete mystackslug mydsid
```


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
