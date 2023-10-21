package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/spf13/cobra"
)

var domain string

var createProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Creates a whole REST Project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("The name of the domain is required.")
		}
		if modName == "" {
			fmt.Println("The name of the module is required.")
		}
		
		d := types.NewDomainTmpl(directory, domain, modName, "templates", false, types.StandardProject)	

		project := rest_cli.ProjectTypeFactory(d)
		err := project.Create()
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")

	rootCmd.AddCommand(createProjectCmd)
}
