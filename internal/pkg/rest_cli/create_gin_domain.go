package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type ginDomain struct {
	DomainTmpl *types.DomainTmpl
}

func (gd *ginDomain) Create() error {
	err := util.CreateFileSkeleton(gd.createHandler())

	err = util.CreateFileSkeleton(gd.createService())
	err = util.CreateFileSkeleton(gd.createStorage())
	err = util.CreateFileSkeleton(gd.createModel())

	return err
}

func NewGinDomain(d *types.DomainTmpl) *ginDomain {
	return &ginDomain{
		DomainTmpl: d,
	}
}

func (gd *ginDomain) createHandler() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gd.DomainTmpl.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", gd.DomainTmpl.Domain, "handler")

	handlerFileInput := util.CreateFileInput(gd.DomainTmpl, "templates/gin/handler_tmpl.txt", filePath, handlerFileName)
	handlerInterfaceFileInput := util.CreateFileInput(gd.DomainTmpl, "templates/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput}
}
func (gd *ginDomain) createService() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gd.DomainTmpl.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", gd.DomainTmpl.Domain, "service")

	serviceFileInput := util.CreateFileInput(gd.DomainTmpl, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := util.CreateFileInput(gd.DomainTmpl, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}
func (gd *ginDomain) createStorage() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gd.DomainTmpl.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", gd.DomainTmpl.Domain, "storage")

	storageFileInput := util.CreateFileInput(
		gd.DomainTmpl,
		fmt.Sprintf("%s/storage_tmpl.txt", gd.DomainTmpl.TemplatePath),
		filePath,
		storageFileName,
	)
	storageInterfaceFileInput := util.CreateFileInput(
		gd.DomainTmpl,
		fmt.Sprintf("%s/storage_interface_tmpl.txt", gd.DomainTmpl.TemplatePath),
		filePath,
		"storage.go",
	)

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput}
}
func (gd *ginDomain) createModel() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gd.DomainTmpl.Domain, "models")
	modelFileName := fmt.Sprintf("%s_%s.go", gd.DomainTmpl.Domain, "models")

	modelFileInput := util.CreateFileInput(gd.DomainTmpl, "templates/gin/model_tmpl.txt", filePath, modelFileName)

	return []*types.FileInputs{modelFileInput}
}
