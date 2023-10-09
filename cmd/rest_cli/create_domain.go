package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-template-cli/pkg/rest_cli"
	"github.com/spf13/cobra"
)

var name string
var includeUtils bool = false

var createDomainCmd = &cobra.Command{
	Use:     "domain",
	Aliases: []string{"d"},
	Short:   "Creates a whole domain from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args", args)
		res := rest_cli.CreateDomainFromScratch(name, modName, directory, includeUtils)
		fmt.Println(res)
	},
}

func init() {
	createDomainCmd.Flags().StringVarP(&name, "name", "n", "test", "The name of the domain")
	createDomainCmd.Flags().StringVarP(&modName, "goModName", "m", "example.com/test", "The Go mod name")
	createDomainCmd.Flags().StringVarP(&directory, "directory", "p", "test-dir", "The Directory to create the project")
	createDomainCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")
	rootCmd.AddCommand(createDomainCmd)
}
