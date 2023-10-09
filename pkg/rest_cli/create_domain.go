package rest_cli

import (
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-template-cli/types"
	"github.com/Cheveo/go-rest-template-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateDomainFromScratch(domain, modName, directory string, includeUtils bool) error {
	var resolvedUserPath string
	if directory != "" {
		userPath, _ := os.UserHomeDir()
		resolvedUserPath = userPath
	} else {
		resolvedUserPath = ""
	}

	d := types.DomainTmpl{
		Directory:         filepath.Join(resolvedUserPath, directory),
		Domain:            domain,
		CapitalizedDomain: cases.Title(language.English, cases.Compact).String(domain),
		GoMod:             modName,
	}

	err := util.CreateFileSkeleton(util.CreateHandler(&d))

	err = util.CreateFileSkeleton(util.CreateService(&d))
	err = util.CreateFileSkeleton(util.CreateStorage(&d))
	err = util.CreateFileSkeleton(util.CreateModel(&d))
	
    if includeUtils {
		err = util.CreateFileSkeleton(util.CreateUtil(&d))
	}

	return err
}
