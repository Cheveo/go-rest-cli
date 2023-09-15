package rest_cli

import (
	"fmt"

	"cheveo.de/Development/go-rest-cli/pkg/rest_cli"
	"github.com/spf13/cobra"
)

var domain, modName, directory string
var createDomainCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Creates a whole domain from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("args", args)
		res := rest_cli.CreateDomainFromScratch(domain, modName, directory)
		fmt.Println(res)
	},
}

func init() {
	createDomainCmd.Flags().StringVarP(&domain, "domain", "d", "test", "The name of the domain")
	createDomainCmd.Flags().StringVarP(&modName, "goModName", "m", "example.com/test", "The Go mod name")
	createDomainCmd.Flags().StringVarP(&directory, "directory", "p", "test-dir", "The Directory to create the project")

	rootCmd.AddCommand(createDomainCmd)
}
