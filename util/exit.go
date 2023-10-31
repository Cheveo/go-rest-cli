package util

import "os"

func Exit(message string, code int) {
	PrintError(message)
	os.Exit(code)
}
