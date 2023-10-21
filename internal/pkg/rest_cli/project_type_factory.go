package rest_cli

import "github.com/Cheveo/go-rest-cli/types"

func ProjectTypeFactory(d *types.DomainTmpl) Create {
	switch d.Type {
	case types.StandardProject:
		return NewStandardProject(d)
	case types.StandardDomain:
		return NewStandardDomain(d)
	case types.GinDomain:
		return NewGinDomain(d)
	case types.GinProject:
		return NewGinProject(d)
	}

	return nil
}
