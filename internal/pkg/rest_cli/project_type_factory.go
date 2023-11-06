package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/types"
)

func ProjectTypeFactory(d *types.DomainTmpl) (Creator, error) {
	switch d.Type {
	case types.StandardProject:
		return NewStandardProject(d), nil
	case types.StandardDomain:
		return NewStandardDomain(d), nil
	case types.GinProject:
		return NewGinProject(d), nil
	case types.GinDomain:
		return NewGinDomain(d), nil
	}

	return nil, fmt.Errorf("Project type does not exist: %s", d.Type)
}
