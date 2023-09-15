package util

import (
	"os"
)

func MakeDirAtPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)

		return err
	}

	return nil
}
