package main

import (
	"cheveo.de/Development/go-rest-cli/cmd"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	Create = "create"
)

func main() {
	fmt.Println(os.Args[1:])

	if len(os.Args) < 2 {
		fmt.Println("Please pass one of the following commands: create")
		os.Exit(0)
	}

	command := os.Args[1]

	fmt.Println(command)
	switch command {
	case Create:
		fmt.Println("Hello create")
		cmd.CreateDomainFromScratch()
	}
}
