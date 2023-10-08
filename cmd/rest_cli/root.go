package rest_cli

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var modName, directory string

var rootCmd = &cobra.Command{
    Use:  "rest-cli",
    Short: "rest-cli - a simple CLI to create go rest apps",
    Long: `rest-cli is an opinionated tool to scaffold rest projects in golang. 
    One can create microservices and monolithic structured Web Servers.`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
