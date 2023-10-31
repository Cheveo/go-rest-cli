package types

import (
	"os"
	"path/filepath"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ProjectType int

const (
	StandardDomain ProjectType = iota
	StandardProject
	GinProject
	GinDomain
)

func (pt ProjectType) String() string {
	switch pt {
	case StandardProject:
		return "StandardProject"
	case StandardDomain:
		return "StandardDomain"
	case GinProject:
		return "GinProject"
	case GinDomain:
		return "GinDomain"
	}

	return "Unknown"
}

type DomainTmpl struct {
	Directory         string
	Domain            string
	CapitalizedDomain string
	StructName        string
	GoMod             string
	Type              ProjectType
	IncludeUtils      bool
	TemplatePath      string
}

func NewDomainTmpl(directory, domain, modName, templatePath string, includeUtils bool, t ProjectType) (*DomainTmpl, error) {
	var resolvedUserPath string
	if directory != "" {
		userPath, _ := os.UserHomeDir()
		resolvedUserPath = userPath
	} else {
		path, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		resolvedUserPath = path
	}

	return &DomainTmpl{
		Directory:         filepath.Join(resolvedUserPath, directory),
		Domain:            domain,
		CapitalizedDomain: cases.Title(language.English, cases.Compact).String(domain),
		GoMod:             modName,
		Type:              t,
		IncludeUtils:      includeUtils,
		TemplatePath:      templatePath,
	}, nil
}
