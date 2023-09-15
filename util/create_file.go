package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(filePath string, name string, dir string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s", filePath, name)
	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return nil, err
	}

	return file, err
}
