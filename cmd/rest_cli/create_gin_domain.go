package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createGinDomainCmd = &cobra.Command{
	Use:     "gin-domain",
	Aliases: []string{"gd"},
	Short:   "Creates a whole gin domain from scratch",
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

		d := types.NewDomainTmpl(directory, name, modName, "templates/gin", false, types.GinDomain)

		domain := rest_cli.ProjectTypeFactory(d)
		err := domain.Create()

		if err != nil {
			util.Exit(err.Error(), 2)
		}

		color.Set(color.FgGreen)

		fmt.Printf("Successfully created gin domain: %s", name)
	},
}

func init() {
	createGinDomainCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the domain")
	createGinDomainCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createGinDomainCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createGinDomainCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createGinDomainCmd)
}
