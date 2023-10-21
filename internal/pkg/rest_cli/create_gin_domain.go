package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type ginDomain struct {
	Domain *types.DomainTmpl
}

func (s *ginDomain) Create() error {
	err := util.CreateFileSkeleton(createHandler(s.Domain))

	err = util.CreateFileSkeleton(createService(s.Domain))
	err = util.CreateFileSkeleton(createStorage(s.Domain))
	err = util.CreateFileSkeleton(util.CreateModel(s.Domain))

	return err
}

func NewGinDomain(d *types.DomainTmpl) *ginDomain {
	return &ginDomain{
		Domain: d,
	}
}

func createHandler(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "handler")

	handlerFileInput := util.CreateFileInput(domain, "templates/gin/handler_tmpl.txt", filePath, handlerFileName)
	handlerInterfaceFileInput := util.CreateFileInput(domain, "templates/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput}
}
func createService(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "service")

	serviceFileInput := util.CreateFileInput(domain, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := util.CreateFileInput(domain, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}
func createStorage(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "storage")

	storageFileInput := util.CreateFileInput(
		domain,
		fmt.Sprintf("%s/storage_tmpl.txt", domain.TemplatePath),
		filePath,
		storageFileName,
	)
	storageInterfaceFileInput := util.CreateFileInput(
		domain,
		fmt.Sprintf("%s/storage_interface_tmpl.txt", domain.TemplatePath),
		filePath,
		"storage.go",
	)

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput}
}
