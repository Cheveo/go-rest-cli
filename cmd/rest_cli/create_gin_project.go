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

var createGinProjectCmd = &cobra.Command{
	Use:     "gin-project",
	Aliases: []string{"gp"},
	Short:   "Creates a whole gin and gonic powered project from scratch",
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

		d := types.NewDomainTmpl(directory, domain, modName, "templates/gin", false, types.GinProject)
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
			"Successfully created gin webservice project \npath: %s \nwith domain: %s, \nwith module name: %s",
			directory,
			domain,
			modName,
		)
	},
}

func init() {
	createGinProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createGinProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createGinProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createGinProjectCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createGinProjectCmd)
}
