package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/fatih/color"
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
		color.Set(color.FgRed)
		defer color.Unset()

		if name == "" {
			util.Exit("[ERROR] The name of the domain is required.", 2)
		}
		if modName == "" {
			util.Exit("[ERROR] The name of the module is required.", 2)
		}

		d := types.NewDomainTmpl(directory, name, modName, "templates", includeUtils, types.StandardDomain)

		domain := rest_cli.ProjectTypeFactory(d)
		err := domain.Create()

		if err != nil {
			util.Exit(err.Error(), 2)
		}
		color.Set(color.FgGreen)

		fmt.Printf("Successfully created standard domain \nwith domain: %s", name)
	},
}

func init() {
	createDomainCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the domain")
	createDomainCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createDomainCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createDomainCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createDomainCmd)
}
