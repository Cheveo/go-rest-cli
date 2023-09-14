package util

import (
	"fmt"
	"os"
)

func CreateFile(filePath string, name string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s", filePath, name)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, err
}
