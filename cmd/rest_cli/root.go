package rest_cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	directory string
	modName   string
	domain string

	rootCmd = &cobra.Command{
		Use:   "go-rest-cli",
		Short: "go-rest-cli - a simple CLI to create go rest apps",
		Long: `go-rest-cli is an opinionated tool to scaffold rest projects in golang.\nOne can create separate domains or whole webservice projects,\n
	powered either by the standard lib and gorilla mux or gin and gonic and Gorm for Web Servers.`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
