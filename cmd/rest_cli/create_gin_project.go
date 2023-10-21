package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/spf13/cobra"
)

var createGinProjectCmd = &cobra.Command{
	Use:     "gin-project",
	Aliases: []string{"gp"},
	Short:   "Creates a whole gin project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			panic("The name of the domain is required.")
		}
		if modName == "" {
			panic("The name of the module is required.")
		}

		d := types.NewDomainTmpl(directory, name, modName, "templates/gin", false, types.GinProject)

		project := rest_cli.ProjectTypeFactory( d)
		err := project.Create()

		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	createGinProjectCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the domain")
	createGinProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createGinProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createGinProjectCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createGinProjectCmd)
}
