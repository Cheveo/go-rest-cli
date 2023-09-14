package util

import (
	"fmt"
	"os"
	"text/template"

	"cheveo.de/Development/go-rest-cli/types"
)

func CreateFileSkeleton(domain *types.DomainTmpl, typeDescription string) error {

	handlerInterfaceTmplFilePath := fmt.Sprintf("templates/%s_interface_tmpl.txt", typeDescription)
	handlerTmplFilePath := fmt.Sprintf("templates/%s_tmpl.txt", typeDescription)
	handlerInterfaceTmpl := template.Must(template.ParseFiles(handlerInterfaceTmplFilePath))
	handlerTmpl := template.Must(template.ParseFiles(handlerTmplFilePath))

	filePath := fmt.Sprintf("%s/%s", domain.Domain, typeDescription)

	if err := MakeDirAtPath(filePath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	handlerFile, err := CreateFile(filePath, fmt.Sprintf("%s_%s.go", domain.Domain, typeDescription))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer handlerFile.Close()

	handlerInterfaceFile, err := CreateFile(filePath, fmt.Sprintf("%s.go", typeDescription))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer handlerInterfaceFile.Close()

	err = handlerTmpl.Execute(handlerFile, domain)
	err = handlerInterfaceTmpl.Execute(handlerInterfaceFile, domain)

	return err
}
