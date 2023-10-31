package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/config"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type standardProject struct {
	Domain *types.DomainTmpl
}

func (s *standardProject) Create() error {

	var c types.Configuration

	err := config.ReadConfig(&c, "std-project.yaml")
	if err != nil {
		return err
	}

	return util.Create(s.Domain, &c)
}

func NewStandardProject(d *types.DomainTmpl) *standardProject {
	return &standardProject{
		Domain: d,
	}
}
