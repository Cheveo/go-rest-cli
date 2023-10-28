package globals

import (
	"embed"
	"html/template"
	"io/fs"
)

const (
	templatesDir    = "templates"
	ginTemplatesDir = "templates/gin"
)

var (
	//go:embed all:templates
	files     embed.FS
	Templates map[string]*template.Template
)

func LoadTemplates() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	err = aggregateTemplates(tmplFiles, templatesDir)
	if err != nil {
		return err
	}

	ginTmplFiles, err := fs.ReadDir(files, ginTemplatesDir)
	if err != nil {
		return err
	}
	err = aggregateTemplates(ginTmplFiles, ginTemplatesDir)

	return err
}

func aggregateTemplates(tmplFiles []fs.DirEntry, dir string) error {
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

	return nil
}
