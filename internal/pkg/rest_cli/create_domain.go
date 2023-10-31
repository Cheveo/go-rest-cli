package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/config"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type standardDomain struct {
	Domain *types.DomainTmpl
}

func (s *standardDomain) Create() error {
	var c types.Configuration

	err := config.ReadConfig(&c, "std-domain.yaml")
	if err != nil {
		return err
	}

	return util.Create(s.Domain, &c)
}

func NewStandardDomain(d *types.DomainTmpl) *standardDomain {
	return &standardDomain{
		Domain: d,
	}
}
