package rest_cli

import (
	"fmt"
	"os"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:     "std-project",
	Aliases: []string{"sp"},
	Short:   "Creates a whole standard lib powered REST Project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			util.Exit("[ERROR] The name of the domain is required.", 1)
		}
		if mod == "" {
			util.Exit("[ERROR] The name of the module is required.", 1)
		}
		if name == "" {
			util.Exit("[ERROR] The name of the project is required.", 1)
		}

		d, err := types.NewDomainTmpl(directory, domain, mod, name, types.StandardProject)
		if err != nil {
			util.Exit(err.Error(), 1)
		}
		if _, err := os.Stat(d.Directory); !os.IsNotExist(err) {
			util.Exit("[ERROR] Directory exists", 1)
		}

		project := rest_cli.ProjectTypeFactory(d)
		err = project.Create()
		if err != nil {
			util.Exit(err.Error(), 1)
		}

		util.PrintSuccess(fmt.Sprintf(
			"Successfully created standard lib webservice project \npath: %s \nwith domain: %s, \nwith module name: %s",
			directory,
			domain,
			mod,
		))
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&mod, "goModName", "m", "", "The Go mod name")
	createProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")

	rootCmd.AddCommand(createProjectCmd)
}
