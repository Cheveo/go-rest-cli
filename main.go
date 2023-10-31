package main

import (
	"github.com/Cheveo/go-rest-cli/cmd/rest_cli"
	"github.com/Cheveo/go-rest-cli/globals"
)

func main() {
	globals.LoadTemplates()
	rest_cli.Execute()
}

