package types

type CommandInput struct {
	Required    bool
	Input       string
	Description string
}

type DomainTmpl struct {
	Domain            string
	CapitalizedDomain string
	StructName        string
	GoMod             string
}
