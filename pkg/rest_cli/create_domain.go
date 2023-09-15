package rest_cli 

import (
	"fmt"
	"os"
	"path/filepath"

	"cheveo.de/Development/go-rest-cli/types"
	"cheveo.de/Development/go-rest-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateDomainFromScratch(domain, modName, directory string) error {
	var resolvedUserPath string
	if directory != "" {
		userPath, _ := os.UserHomeDir()
		resolvedUserPath = userPath
	} else {
		resolvedUserPath = ""
	}

	d := types.DomainTmpl{
		Directory:         filepath.Join(resolvedUserPath, directory),
		Domain:            domain,
		CapitalizedDomain: cases.Title(language.English, cases.Compact).String(domain),
		GoMod:             modName,
	}

	absolutePath, _ := os.UserHomeDir()
	fmt.Println(fmt.Sprintf("PATH %s", absolutePath))
	fmt.Println("Domain", d.Domain)
	fmt.Println("Module Name", modName)

	err := util.CreateFileSkeleton(createMain(&d))
	err = util.CreateFileSkeleton(createGoMod(&d))
	err = util.CreateFileSkeleton(createHandler(&d))
	err = util.CreateFileSkeleton(createUtil(&d))
	err = util.CreateFileSkeleton(createService(&d))
	err = util.CreateFileSkeleton(createStorage(&d))
	err = util.CreateFileSkeleton(createModel(&d))
	err = util.CreateFileSkeleton(createServer(&d))

	return err
}

func createFileInput(domain *types.DomainTmpl, templatePath string, filePath string, modelFileName string) *types.FileInputs {
	fileInput := types.FileInputs{
		TemplatePath: templatePath,
		FilePath:     filePath,
		FileName:     modelFileName,
		Context:      domain,
	}
	return &fileInput
}

func createGoMod(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/go.mod.txt",
		FilePath:     "",
		FileName:     "go.mod",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}
}
func createMain(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/main.txt",
		FilePath:     "",
		FileName:     "main.go",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}
}
func createServer(domain *types.DomainTmpl) []*types.FileInputs {
	serverFileInput := types.FileInputs{
		TemplatePath: "templates/server.txt",
		FilePath:     "server",
		FileName:     "server.go",
		Context:      domain,
	}

	return []*types.FileInputs{&serverFileInput}

}
func createModel(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "models")
	modelFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "models")

	modelFileInput := createFileInput(domain, "templates/model_tmpl.txt", filePath, modelFileName)
	modelInterfaceFileInput := createFileInput(domain, "templates/model_interface_tmpl.txt", filePath, "models.go")

	return []*types.FileInputs{modelFileInput, modelInterfaceFileInput}
}
func createStorage(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "storage")
	storageSqlStatementsFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "sql_statements")

	storageFileInput := createFileInput(domain, "templates/storage_tmpl.txt", filePath, storageFileName)
	storageSqlStatementsFileInput := createFileInput(domain, "templates/storage_sql_statements_tmpl.txt", filePath, storageSqlStatementsFileName)
	storageInterfaceFileInput := createFileInput(domain, "templates/storage_interface_tmpl.txt", filePath, "storage.go")

	return []*types.FileInputs{storageFileInput, storageInterfaceFileInput, storageSqlStatementsFileInput}
}

func createService(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "service")

	serviceFileInput := createFileInput(domain, "templates/service_tmpl.txt", filePath, serviceFileName)
	serviceInterfaceFileInput := createFileInput(domain, "templates/service_interface_tmpl.txt", filePath, "service.go")

	return []*types.FileInputs{serviceFileInput, serviceInterfaceFileInput}
}

func createHandler(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "handler")

	handlerFileInput := createFileInput(domain, "templates/handler_tmpl.txt", filePath, handlerFileName)
	handlerHttpErrorFileInput := createFileInput(domain, "templates/handler_http_errors_tmpl.txt", filePath, "handle_http_errors.go")
	handlerInterfaceFileInput := createFileInput(domain, "templates/handler_interface_tmpl.txt", filePath, "handler.go")

	return []*types.FileInputs{handlerFileInput, handlerInterfaceFileInput, handlerHttpErrorFileInput}
}

func createUtil(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := "util"

	makeHttpHandlerFileInput := createFileInput(domain, "templates/util_make_http_handler.txt", filePath, "make_http_handler.go")
	writeJsonFileInput := createFileInput(domain, "templates/util_write_json.txt", filePath, "write_json.go")

	return []*types.FileInputs{makeHttpHandlerFileInput, writeJsonFileInput}
}
