package botbuilder

// Subcommand represents a Discord subcommand.
type Subcommand struct {
	Name           string
	Description    string
	Aliases        []string
	ExecuteCommand bool
	Execute        CommandFunction
}
