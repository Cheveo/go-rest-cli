package rest_cli

import (
	"fmt"
	"os"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/spf13/cobra"
)

var createGinProjectCmd = &cobra.Command{
	Use:     "gin-project",
	Aliases: []string{"gp"},
	Short:   "Creates a whole gin and gonic powered project from scratch",
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

		d, err := types.NewDomainTmpl(directory, domain, mod, name, types.GinProject)
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
			"Successfully created gin webservice project \npath: %s \nwith domain: %s, \nwith module name: %s",
			directory,
			domain,
			mod,
		))
	},
}

func init() {
	createGinProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createGinProjectCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the project")
	createGinProjectCmd.Flags().StringVarP(&mod, "module", "m", "", "The Go mod name")
	createGinProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createGinProjectCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createGinProjectCmd)
}
