package util

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Cheveo/go-rest-cli/types"
)

func CreateFileSkeleton(files []*types.FileInputs) error {
	for _, file := range files {
		tmpl := template.Must(template.ParseFiles(filepath.Join(file.TemplatePath)))

		if err := MakeDirAtPath(filepath.Join(file.Context.Directory, file.FilePath)); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Creating File...Path: %s | Filename: %s", file.FilePath, file.FileName))

		createdFile, err := CreateFile(file.FilePath, file.FileName, file.Context.Directory)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer createdFile.Close()

		err = tmpl.Execute(createdFile, file.Context)

		if err != nil {
			return err
		}
	}

	return nil
}
