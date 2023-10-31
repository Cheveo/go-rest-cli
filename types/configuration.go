package types

import "fmt"

type File struct {
	TemplatePath            string `yaml:"templatePath,omitempty"`
	FilePath                string `yaml:"filePath,omitempty"`
	FileName                string `yaml:"fileName,omitempty"`
	HasFilePathDomainPrefix bool   `yaml:"hasFilePathDomainPrefix,omitempty"`
	HasFileNameDomainPrefix bool   `yaml:"hasFileNameDomainPrefix,omitempty"`
}

type Configuration struct {
	Domain string `yaml:"domain,omitempty"`
	Files  []File `yaml:"files,omitempty"`
}

func (f *File) GetFileName(domain string) string {
	if f.HasFileNameDomainPrefix {
		return fmt.Sprintf("%s_%s", domain, f.FileName)
	}

	return f.FileName
}
func (f *File) GetFilePath(domain string) string {
	if f.HasFilePathDomainPrefix {
		return fmt.Sprintf("%s/%s", domain, f.FilePath)
	}

	return f.FilePath
}
