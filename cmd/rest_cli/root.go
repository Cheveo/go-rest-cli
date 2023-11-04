package rest_cli

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	directory string
	mod       string
	domain    string
	name      string
	ascii     = color.New(color.FgBlue).SprintFunc()(`
..######....#######..........########..########..######..########..........######..##.......####
.##....##..##.....##.........##.....##.##.......##....##....##............##....##.##........##.
.##........##.....##.........##.....##.##.......##..........##............##.......##........##.
.##...####.##.....##.#######.########..######....######.....##....#######.##.......##........##.
.##....##..##.....##.........##...##...##.............##....##............##.......##........##.
.##....##..##.....##.........##....##..##.......##....##....##............##....##.##........##.
..######....#######..........##.....##.########..######.....##.............######..########.####
`)

	rootCmd = &cobra.Command{
		Use:   "go-rest-cli",
		Short: "go-rest-cli - a simple CLI to create go rest apps",
		Long: ascii + `
Is an opinionated tool to scaffold rest projects in golang.
One can create separate domains or whole webservice projects, 
powered by:

	💎 Standard Lib + gorilla mux
	💎 Gin & Gonic + Gorm

for Web Servers.
`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	setCobraTemplate()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func setCobraTemplate() {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTemplate)
}
