package command

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Subcommands []Subcommand
}
