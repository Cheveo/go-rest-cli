package types

type FileInput struct {
	TemplatePath string
	FilePath     string
	FileName     string
	Context      *DomainTmpl
}
