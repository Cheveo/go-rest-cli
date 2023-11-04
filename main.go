package main

import (
	"github.com/Cheveo/go-rest-cli/cmd/rest_cli"
	"github.com/Cheveo/go-rest-cli/templates"
	"github.com/Cheveo/go-rest-cli/util"
)

func main() {
	err := templates.LoadTemplates()
	if err != nil {
		util.Exit("Could not load templates", 1)
	}
	rest_cli.Execute()
}

