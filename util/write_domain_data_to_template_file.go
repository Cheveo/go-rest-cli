package util

import (
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/globals"
	"github.com/Cheveo/go-rest-cli/types"
)

func WriteDomainDataToTemplateFile(domain *types.DomainTmpl, configuration *types.Configuration) error {
	for _, file := range configuration.Files {
		filePath := file.GetFilePath(domain.Domain)
		fileName := file.GetFileName(domain.Domain)

		tmpls := globals.Templates

		// templates are organized via template path
		tmpl := tmpls[file.TemplatePath]

		if err := MakeDirAtPath(filepath.Join(domain.Directory, filePath)); err != nil {
			return err
		}

		createdFile, err := CreateFile(filePath, fileName, domain.Directory)
		if err != nil {
			return err
		}
		defer createdFile.Close()

		err = tmpl.Execute(createdFile, domain)
		if err != nil {
			return err
		}
	}

	return nil
}
