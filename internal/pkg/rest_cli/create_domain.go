package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
)

type standardDomain struct {
	Domain *types.DomainTmpl
}

func (s *standardDomain) Create() error {
	err := util.CreateFileSkeleton(s.createHandler())

	err = util.CreateFileSkeleton(s.createService())
	err = util.CreateFileSkeleton(s.createStorage())
	err = util.CreateFileSkeleton(s.createModel())

	if s.Domain.IncludeUtils {
		err = util.CreateFileSkeleton(s.createUtil())
	}

	return err
}

func NewStandardDomain(d *types.DomainTmpl) *standardDomain {
	return &standardDomain{
		Domain: d,
	}
}

func (s *standardDomain) createService() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", s.Domain.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", s.Domain.Domain, "service")

	serviceFileInput := util.CreateFileInput(s.Domain, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := util.CreateFileInput(s.Domain, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}
func (s *standardDomain) createModel() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", s.Domain.Domain, "models")
	modelFileName := fmt.Sprintf("%s_%s.go", s.Domain.Domain, "models")

	modelFileInput := util.CreateFileInput(s.Domain, "templates/model_tmpl.txt", filePath, modelFileName)
	modelInterfaceFileInput := util.CreateFileInput(s.Domain, "templates/model_interface_tmpl.txt", filePath, "models.go")

	return []*types.FileInputs{modelFileInput, modelInterfaceFileInput}
}
func (s *standardDomain) createStorage() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", s.Domain.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", s.Domain.Domain, "storage")
	storageSqlStatementsFileName := fmt.Sprintf("%s_%s.go", s.Domain.Domain, "sql_statements")

	storageFileInput := util.CreateFileInput(s.Domain, "templates/storage_tmpl.txt", filePath, storageFileName)
	storageSqlStatementsFileInput := util.CreateFileInput(s.Domain, "templates/storage_sql_statements_tmpl.txt", filePath, storageSqlStatementsFileName)
	storageInterfaceFileInput := util.CreateFileInput(s.Domain, "templates/storage_interface_tmpl.txt", filePath, "storage.go")

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput, storageSqlStatementsFileInput}
}
func (s *standardDomain) createHandler() []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", s.Domain.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", s.Domain.Domain, "handler")

	handlerFileInput := util.CreateFileInput(s.Domain, "templates/handler_tmpl.txt", filePath, handlerFileName)
	handlerHttpErrorFileInput := util.CreateFileInput(s.Domain, "templates/handler_http_errors_tmpl.txt", filePath, "handle_http_errors.go")
	handlerInterfaceFileInput := util.CreateFileInput(s.Domain, "templates/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput, handlerHttpErrorFileInput}
}
func (s *standardDomain) createUtil() []*types.FileInputs {
	filePath := "util"

	makeHttpHandlerFileInput := util.CreateFileInput(s.Domain, "templates/util_make_http_handler.txt", filePath, "make_http_handler.go")
	writeJsonFileInput := util.CreateFileInput(s.Domain, "templates/util_write_json.txt", filePath, "write_json.go")

	return []*types.FileInputs{makeHttpHandlerFileInput, writeJsonFileInput}
}
