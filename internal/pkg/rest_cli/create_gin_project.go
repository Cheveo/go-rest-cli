package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/config"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type ginProject struct {
	DomainTmpl *types.DomainTmpl
}

func (gp *ginProject) Create() error {
	var c types.Configuration

	err := config.ReadConfig(&c, "gin-project.yaml")
	if err != nil{
		return err
	}

	return util.WriteDomainDataToTemplateFile(gp.DomainTmpl, &c)
}

func NewGinProject(d *types.DomainTmpl) *ginProject {
	return &ginProject{
		DomainTmpl: d,
	}
}
