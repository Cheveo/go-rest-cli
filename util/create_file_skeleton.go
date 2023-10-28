package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/globals"
	"github.com/Cheveo/go-rest-cli/types"
)

func CreateFileSkeleton(files []*types.FileInputs) error {
	tmpls := globals.Templates
	for _, file := range files {

		// templates are organized via template path
		tmpl := tmpls[file.TemplatePath]

		if err := MakeDirAtPath(filepath.Join(file.Context.Directory, file.FilePath)); err != nil {
			fmt.Println("error", err.Error())
			os.Exit(1)
		}


		createdFile, err := CreateFile(file.FilePath, file.FileName, file.Context.Directory)
		if err != nil {
			fmt.Println("error", err.Error())
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
