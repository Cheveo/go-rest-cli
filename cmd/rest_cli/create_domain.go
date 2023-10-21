package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types" 
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
		if name == "" {
			fmt.Println("The name of the domain is required.")
		}
		if modName == "" {
			fmt.Println("The name of the module is required.")
		}

		d := types.NewDomainTmpl(directory, name, modName, "templates", includeUtils, types.StandardDomain)	

		domain := rest_cli.ProjectTypeFactory(d)
		err := domain.Create()

		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	createDomainCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the domain")
	createDomainCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createDomainCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createDomainCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createDomainCmd)
}


