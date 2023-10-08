package rest_cli

import (
	"fmt"

	"cheveo.de/Development/go-rest-cli/pkg/rest_cli"
	"github.com/spf13/cobra"
)

var domain string

var createProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Creates a whole REST Project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("args", args)
		res := rest_cli.CreateProjectFromScratch(domain, modName, directory)
		fmt.Println(res)
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&domain, "domain", "d", "test", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "example.com/test", "The Go mod name")
	createProjectCmd.Flags().StringVarP(&directory, "directory", "p", "test-dir", "The Directory to create the project")

	rootCmd.AddCommand(createProjectCmd)
}
