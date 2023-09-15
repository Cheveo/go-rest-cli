package types

type FileInputs struct {
	TemplatePath string
	FilePath     string
	FileName     string
	Context      *DomainTmpl
}
