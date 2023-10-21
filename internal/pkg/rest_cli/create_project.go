package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type standardProject struct{
	Domain *types.DomainTmpl
}

func (s *standardProject) Create() error {
	err := util.CreateFileSkeleton(util.CreateMain(s.Domain))
	err = util.CreateFileSkeleton(util.CreateGoMod(s.Domain))
	err = util.CreateFileSkeleton(util.CreateHandler(s.Domain))
	err = util.CreateFileSkeleton(util.CreateUtil(s.Domain))
	err = util.CreateFileSkeleton(util.CreateService(s.Domain))
	err = util.CreateFileSkeleton(util.CreateStorage(s.Domain))
	err = util.CreateFileSkeleton(util.CreateModel(s.Domain))
	err = util.CreateFileSkeleton(util.CreateServer(s.Domain))

	return err
}

func NewStandardProject(d *types.DomainTmpl) *standardProject {
	return &standardProject{
		Domain: d,
	}
}
