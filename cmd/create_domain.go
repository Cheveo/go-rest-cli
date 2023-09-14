package cmd

import (
	"flag"
	"fmt"
	"os"

	"cheveo.de/Development/go-rest-cli/types"
	"cheveo.de/Development/go-rest-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateDomainFromScratch() error {
	initCmd := flag.NewFlagSet("create", flag.ExitOnError)

	var d, modName string

	initCmd.StringVar(&d, "domain", "", "The name of the domain")
	initCmd.StringVar(&modName, "goModName", "", "The Go mod name")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Add the create command")
		os.Exit(0)
	}

	initCmd.Parse(os.Args[2:])

	domain := types.DomainTmpl{
		Domain:            d,
		CapitalizedDomain: cases.Title(language.English, cases.Compact).String(d),
		GoMod:             modName,
	}
	fmt.Println(d)
	fmt.Println(modName)

	err := util.CreateFileSkeleton(&domain, "handler")
	err = util.CreateFileSkeleton(&domain, "service")
	err = util.CreateFileSkeleton(&domain, "storage")
	err = util.CreateFileSkeleton(&domain, "model")

	return err
}
