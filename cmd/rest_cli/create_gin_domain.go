package rest_cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/internal/pkg/rest_cli"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"github.com/spf13/cobra"
)

var createGinDomainCmd = &cobra.Command{
	Use:     "gin-domain",
	Aliases: []string{"gd"},
	Short:   "Creates a whole gin domain from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			util.Exit("[ERROR] The name of the domain is required.", 1)
		}
		if mod == "" {
			util.Exit("[ERROR] The name of the module is required.", 1)
		}

		domainTmpl, err := types.NewDomainTmpl(directory, domain, mod, "", types.GinDomain)
		if err != nil {
			util.Exit(err.Error(), 1)
		}

		if _, err := os.Stat(filepath.Join(domainTmpl.Directory, domain)); !os.IsNotExist(err) {
			util.Exit("[ERROR] Domain exists", 1)
		}

		d, err := rest_cli.ProjectTypeFactory(domainTmpl)
		if err != nil {
			util.Exit(err.Error(), 1)
		}

		err = d.Create()
		if err != nil {
			util.Exit(err.Error(), 1)
		}

		util.PrintSuccess(fmt.Sprintf("Successfully created gin domain: %s", domain))
	},
}

func init() {
	createGinDomainCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createGinDomainCmd.Flags().StringVarP(&mod, "goModName", "m", "", "The Go mod name")
	createGinDomainCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")
	createGinDomainCmd.Flags().BoolVarP(&includeUtils, "includeUtils", "i", false, "Create with utils: makeHttpHandler and writeJson")

	rootCmd.AddCommand(createGinDomainCmd)
}
