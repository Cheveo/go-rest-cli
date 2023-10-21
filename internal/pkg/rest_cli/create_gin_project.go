package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type ginProject struct {
	DomainTmpl *types.DomainTmpl
}

func (gp *ginProject) Create() error {
	err := util.CreateFileSkeleton(gp.createHandler())

	err = util.CreateFileSkeleton(gp.createService())
	err = util.CreateFileSkeleton(gp.createStorage())
	err = util.CreateFileSkeleton(gp.createDatabase())
	err = util.CreateFileSkeleton(gp.createMain())
	err = util.CreateFileSkeleton(gp.createModel())
	err = util.CreateFileSkeleton(gp.createResponses())
	err = util.CreateFileSkeleton(gp.createErrors())
	err = util.CreateFileSkeleton(gp.createMiddlewares())
	err = util.CreateFileSkeleton(util.CreateGoMod(gp.DomainTmpl))

	return err
}

func NewGinProject(d *types.DomainTmpl) *ginProject {
	return &ginProject{
		DomainTmpl: d,
	}
}

func (gp *ginProject) createHandler() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gp.DomainTmpl.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", gp.DomainTmpl.Domain, "handler")

	handlerFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/handler_tmpl.txt", filePath, handlerFileName)
	handlerInterfaceFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput}
}
func (gp *ginProject) createService() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gp.DomainTmpl.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", gp.DomainTmpl.Domain, "service")

	serviceFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}
func (gp *ginProject) createStorage() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gp.DomainTmpl.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", gp.DomainTmpl.Domain, "storage")

	storageFileInput := util.CreateFileInput(
		gp.DomainTmpl,
		fmt.Sprintf("%s/storage_tmpl.txt", gp.DomainTmpl.TemplatePath),
		filePath,
		storageFileName,
	)
	storageInterfaceFileInput := util.CreateFileInput(
		gp.DomainTmpl,
		fmt.Sprintf("%s/storage_interface_tmpl.txt", gp.DomainTmpl.TemplatePath),
		filePath,
		"storage.go",
	)

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput}
}
func (gp *ginProject) createDatabase() []*types.FileInputs {
	filePath := "db"
	dbFileName := "gorm_db.go"

	gormDbFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/gorm_db_tmpl.txt", filePath, dbFileName)
	dbInterfaceFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/db_tmpl.txt", filePath, "db.go")

	return []*types.FileInputs{gormDbFileInput, dbInterfaceFileInput}
}
func (gp *ginProject) createModel() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", gp.DomainTmpl.Domain, "models")
	modelFileName := fmt.Sprintf("%s_%s.go", gp.DomainTmpl.Domain, "models")

	modelFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/model_tmpl.txt", filePath, modelFileName)

	return []*types.FileInputs{modelFileInput}
}
func (gp *ginProject) createMain() []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/gin/main.txt",
		FilePath:     "cmd/api",
		FileName:     "main.go",
		Context:      gp.DomainTmpl,
	}

	return []*types.FileInputs{&serverFileInput}
}
func (gp *ginProject) createResponses() []*types.FileInputs {
	responseFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/http_responses_tmpl.txt", "responses", "http_response.go")

	return []*types.FileInputs{responseFileInput}
}
func (gp *ginProject) createErrors() []*types.FileInputs {
	errorFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/errors_tmpl.txt", "errors", "http_error.go")

	return []*types.FileInputs{errorFileInput}
}
func (gp *ginProject) createMiddlewares() []*types.FileInputs {
	errorHandlerFileInput := util.CreateFileInput(gp.DomainTmpl, "templates/gin/error_handler_tmpl.txt", "middlewares", "error_handler.go")

	return []*types.FileInputs{errorHandlerFileInput}
}
