package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/pkg/version"
)

// NewCmdVersion returns a new cobra command.
func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  "Show version.",
		Run:   run,
	}

	return cmd
}

// PrintVersion prints out the version.
func PrintVersion() {
	fmt.Println(version.Version)
}

// GetVersion returns the version.
func GetVersion() string {
	return version.Version
}

// run runs the command's action.
func run(cmd *cobra.Command, args []string) {
	PrintVersion()
}
