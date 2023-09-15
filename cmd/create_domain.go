package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cheveo.de/Development/go-rest-cli/types"
	"cheveo.de/Development/go-rest-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateDomainFromScratch() error {
	initCmd := flag.NewFlagSet("create", flag.ExitOnError)

	var domain, modName, directory string

	initCmd.StringVar(&domain, "domain", "", "The name of the domain")
	initCmd.StringVar(&modName, "goModName", "", "The Go mod name")
	initCmd.StringVar(&directory, "directory", "", "The Directory to create the project")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Add the create command")
		os.Exit(0)
	}

	initCmd.Parse(os.Args[2:])

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
	fmt.Println(d)
	fmt.Println(modName)

	err := util.CreateFileSkeleton(createHandler(&d))
	err = util.CreateFileSkeleton(createUtil(&d))
	err = util.CreateFileSkeleton(createService(&d))
	err = util.CreateFileSkeleton(createStorage(&d))
	err = util.CreateFileSkeleton(createModel(&d))
	err = util.CreateFileSkeleton(createServer(&d))

	return err
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
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "model")
	modelFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "model")

	modelFileInput := types.FileInputs{
		TemplatePath: "templates/service_tmpl.txt",
		FilePath:     filePath,
		FileName:     modelFileName,
		Context:      domain,
	}

	modelInterfaceFileInput := types.FileInputs{
		TemplatePath: "templates/service_interface_tmpl.txt",
		FilePath:     filePath,
		FileName:     "model.go",
		Context:      domain,
	}

	return []*types.FileInputs{&modelFileInput, &modelInterfaceFileInput}
}
func createStorage(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "storage")
	storageFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "storage")

	storageFileInput := types.FileInputs{
		TemplatePath: "templates/service_tmpl.txt",
		FilePath:     filePath,
		FileName:     storageFileName,
		Context:      domain,
	}

	storageInterfaceFileInput := types.FileInputs{
		TemplatePath: "templates/service_interface_tmpl.txt",
		FilePath:     filePath,
		FileName:     "storage.go",
		Context:      domain,
	}

	return []*types.FileInputs{&storageFileInput, &storageInterfaceFileInput}
}

func createService(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "service")
	serviceFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "service")

	serviceFileInput := types.FileInputs{
		TemplatePath: "templates/service_tmpl.txt",
		FilePath:     filePath,
		FileName:     serviceFileName,
		Context:      domain,
	}

	serviceInterfaceFileInput := types.FileInputs{
		TemplatePath: "templates/service_interface_tmpl.txt",
		FilePath:     filePath,
		FileName:     "service.go",
		Context:      domain,
	}

	return []*types.FileInputs{&serviceFileInput, &serviceInterfaceFileInput}
}
func createHandler(domain *types.DomainTmpl) []*types.FileInputs {
	filePath := fmt.Sprintf("%s/%s", domain.Domain, "handler")
	handlerFileName := fmt.Sprintf("%s_%s.go", domain.Domain, "handler")

	handlerFileInput := types.FileInputs{
		TemplatePath: "templates/handler_tmpl.txt",
		FilePath:     filePath,
		FileName:     handlerFileName,
		Context:      domain,
	}

	handlerInterfaceFileInput := types.FileInputs{
		TemplatePath: "templates/handler_interface_tmpl.txt",
		FilePath:     filePath,
		FileName:     "handler.go",
		Context:      domain,
	}

	return []*types.FileInputs{&handlerFileInput, &handlerInterfaceFileInput}
}

func createUtil(domain *types.DomainTmpl) []*types.FileInputs {

	makeHttpHandlerFileInput := types.FileInputs{
		TemplatePath: "templates/util_make_http_handler.txt",
		FilePath:     "util",
		FileName:     "make_http_handler.go",
		Context:      domain,
	}

	writeJsonFileInput := types.FileInputs{
		TemplatePath: "templates/util_write_json.txt",
		FilePath:     "util",
		FileName:     "write_json.go",
		Context:      domain,
	}

	return []*types.FileInputs{&makeHttpHandlerFileInput, &writeJsonFileInput}
}
