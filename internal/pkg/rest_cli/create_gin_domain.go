package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/config"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type ginDomain struct {
	DomainTmpl *types.DomainTmpl
}

func (gd *ginDomain) Create() error {
	var c types.Configuration

	err := config.ReadConfig(&c, "gin-domain.yaml")
	if err != nil {
		return err
	}

	return util.Create(gd.DomainTmpl, &c)
}

func NewGinDomain(d *types.DomainTmpl) *ginDomain {
	return &ginDomain{
		DomainTmpl: d,
	}
}
