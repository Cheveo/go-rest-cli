package rest_cli

import (
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type standardDomain struct{}

func (s *standardDomain) Create(domain, modName, directory string, includeUtils bool) error {
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

func NewStandardDomain() *standardDomain {
	return &standardDomain{}
}
