package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/types"
)

func CreateFile(filePath string, name string, dir string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s", filePath, name)
	file, err := os.Create(filepath.Join(dir, fileName))
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

