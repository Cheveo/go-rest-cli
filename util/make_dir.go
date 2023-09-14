package util

import (
	"os"
	"path/filepath"
)

func MakeDirAtPath(path string) error {
	newpath := filepath.Join(".", path)
	err := os.MkdirAll(newpath, os.ModePerm)

	return err
}
