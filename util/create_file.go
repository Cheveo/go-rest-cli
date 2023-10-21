package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/types"
)

func CreateFile(filePath string, name string, dir string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s/%s", dir, filePath, name)
	file, err := os.Create(filepath.Join(fileName))
	if err != nil {
		return nil, err
	}

	return file, err
}

func CreateFileInput(domain *types.DomainTmpl, templatePath string, filePath string, modelFileName string) *types.FileInputs {
	fileInput := types.FileInputs{
		TemplatePath: templatePath,
		FilePath:     filePath,
		FileName:     modelFileName,
		Context:      domain,
	}
	return &fileInput
}

func CreateGoMod(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/go.mod.txt",
		FilePath:     "",
		FileName:     "go.mod",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}
}
func CreateMain(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/main.txt",
		FilePath:     "",
		FileName:     "main.go",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}
}
func CreateServer(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/server.txt",
		FilePath:     "server",
		FileName:     "server.go",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}

}
func CreateModel(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "models")
	modelFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "models")

	modelFileInput := CreateFileInput(domain, "templates/model_tmpl.txt", filePath, modelFileName)
	modelInterfaceFileInput := CreateFileInput(domain, "templates/model_interface_tmpl.txt", filePath, "models.go")

	return []*types.FileInputs{modelFileInput, modelInterfaceFileInput}
}
func CreateStorage(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "storage")
	storageSqlStatementsFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "sql_statements")

	storageFileInput := CreateFileInput(
		domain,
		fmt.Sprintf("%s/storage_tmpl.txt", domain.TemplatePath),
		filePath,
		storageFileName,
	)
	storageSqlStatementsFileInput := CreateFileInput(
		domain,
		fmt.Sprintf("%s/storage_sql_statements_tmpl.txt", domain.TemplatePath),
		filePath, storageSqlStatementsFileName,
	)
	storageInterfaceFileInput := CreateFileInput(
		domain,
		fmt.Sprintf("%s/storage_interface_tmpl.txt", domain.TemplatePath),
		filePath,
		"storage.go",
	)

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput, storageSqlStatementsFileInput}
}

func CreateService(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "service")

	serviceFileInput := CreateFileInput(domain, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := CreateFileInput(domain, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}

func CreateHandler(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "handler")

	handlerFileInput := CreateFileInput(domain, "templates/handler_tmpl.txt", filePath, handlerFileName)
	handlerHttpErrorFileInput := CreateFileInput(domain, "templates/handler_http_errors_tmpl.txt", filePath, "handle_http_errors.go")
	handlerInterfaceFileInput := CreateFileInput(domain, "templates/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput, handlerHttpErrorFileInput}
}

func CreateUtil(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := "util"

	makeHttpHandlerFileInput := CreateFileInput(domain, "templates/util_make_http_handler.txt", filePath, "make_http_handler.go")
	writeJsonFileInput := CreateFileInput(domain, "templates/util_write_json.txt", filePath, "write_json.go")

	return []*types.FileInputs{makeHttpHandlerFileInput, writeJsonFileInput}
}
