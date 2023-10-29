package rest_cli

import (
	"fmt"
	"os"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:     "std-project",
	Aliases: []string{"sp"},
	Short:   "Creates a whole standard lib powered REST Project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		color.Set(color.FgRed)
		defer color.Unset()
		if domain == "" {
			util.Exit("[ERROR] The name of the domain is required.", 2)
		}
		if modName == "" {
			util.Exit("[ERROR] The name of the module is required.", 2)
		}
		if directory == "" {
			util.Exit("[ERROR] The name of the directory is required.", 2)
		}

		d := types.NewDomainTmpl(directory, domain, modName, "templates", false, types.StandardProject)
		if _, err := os.Stat(d.Directory); !os.IsNotExist(err) {
			util.Exit("[ERROR] Directory exists", 2)
		}

		project := rest_cli.ProjectTypeFactory(d)
		err := project.Create()
		if err != nil {
			util.Exit(err.Error(), 2)
		}

		color.Set(color.FgGreen)

		fmt.Printf(
			"Successfully created standard lib webservice project \npath: %s \nwith domain: %s, \nwith module name: %s",
			directory, 
			domain, 
			modName,
		)
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")

	rootCmd.AddCommand(createProjectCmd)
}
