package types

type CommandInput struct {
	Required    bool
	Input       string
	Description string
}

type DomainTmpl struct {
	Directory         string
	Domain            string
	CapitalizedDomain string
	StructName        string
	GoMod             string
}
