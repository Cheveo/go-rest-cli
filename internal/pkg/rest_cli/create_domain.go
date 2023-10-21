package rest_cli

import (
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type standardDomain struct {
	Domain *types.DomainTmpl
}

func (s *standardDomain) Create() error {
	err := util.CreateFileSkeleton(util.CreateHandler(s.Domain))

	err = util.CreateFileSkeleton(util.CreateService(s.Domain))
	err = util.CreateFileSkeleton(util.CreateStorage(s.Domain))
	err = util.CreateFileSkeleton(util.CreateModel(s.Domain))

	if s.Domain.IncludeUtils {
		err = util.CreateFileSkeleton(util.CreateUtil(s.Domain))
	}

	return err
}

func NewStandardDomain(d *types.DomainTmpl) *standardDomain {
	return &standardDomain{
		Domain: d,
	}
}
