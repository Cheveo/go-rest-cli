package templates

import (
	"embed"
	"html/template"
	"io/fs"
)

const (
	baseTemplatesDir = "base"
	stdTemplatesDir  = "std"
	ginTemplatesDir  = "gin"
)

// It is necessary to embed all needed template 
// files with embed for compilation,
// so they won't be missing in the resulting binary
var (
	//go:embed all:gin
	//go:embed all:std
	//go:embed all:base
	files        embed.FS
	Templates    map[string]*template.Template
	templateDirs = []string{baseTemplatesDir, stdTemplatesDir, ginTemplatesDir}
)

func LoadTemplates() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}

	err := aggregateTemplatesFromTemplateDirectories(templateDirs)
	if err != nil {
		return err
	}

	return nil
}

func aggregateTemplatesFromTemplateDirectories(templateFileDirectories []string) error {
	for _, dir := range templateFileDirectories {
		tmplFiles, err := fs.ReadDir(files, dir)
		if err != nil {
			return err
		}

		for _, tmpl := range tmplFiles {

			if tmpl.IsDir() {
				continue
			}

			pt, err := template.ParseFS(files, dir+"/"+tmpl.Name())
			if err != nil {
				return err
			}

			Templates[dir+"/"+tmpl.Name()] = pt
		}
	}

	return nil
}
