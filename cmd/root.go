package cmd

import (
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/spf13/cobra"

	"github.com/jtyr/gcapi/cmd/apikey"
	"github.com/jtyr/gcapi/cmd/grafana"
	"github.com/jtyr/gcapi/cmd/stack"
	"github.com/jtyr/gcapi/cmd/version"
)

// RootFlags describes a struct that holds flags that can be set on root level of the command.
type RootFlags struct {
	apiToken           string
	apiTokenFile       string
	timestampedLogging bool
	version            bool
}

// flags holds the values of the flags.
var flags = RootFlags{}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "gcapi-cli",
	Short:         "Tool to access Grafana Cloud API",
	Long:          "gcapi-cli is a tool that allows an easy access to the Grafana Cloud API.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Version:       version.GetVersion(),
	Run:           rootRun,
}

// Execute executes the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// GetRootCmd returns the rootCmd.
func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&flags.apiToken, "api-token", "t", "",
		"token used to authenticate to the API")
	rootCmd.PersistentFlags().StringVarP(
		&flags.apiTokenFile, "api-token-file", "f", "",
		"path to a file containing the token used to authenticate to the API")
	rootCmd.PersistentFlags().BoolVar(
		&flags.timestampedLogging, "timestamps", false,
		"enable Log timestamps")

	rootCmd.AddCommand(apikey.NewCmdApiKey())
	rootCmd.AddCommand(grafana.NewCmdGrafana())
	rootCmd.AddCommand(stack.NewCmdStack())
	rootCmd.AddCommand(version.NewCmdVersion())

	// Init
	cobra.OnInitialize(initLogging)
}

// rootRun runs the command's action.
func rootRun(cmd *cobra.Command, args []string) {
	if flags.version {
		version.PrintVersion()
	} else {
		if err := cmd.Usage(); err != nil {
			log.Fatalln(err)
		}
	}
}

// initLogging initializes the logger.
func initLogging() {
	switch logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL")); logLevel {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.SetOutput(ioutil.Discard)
	log.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
			log.TraceLevel,
		},
	})

	formatter := &log.TextFormatter{
		ForceColors: true,
	}

	if flags.timestampedLogging || os.Getenv("LOG_TIMESTAMPS") != "" {
		formatter.FullTimestamp = true
	}

	log.SetFormatter(formatter)
}
