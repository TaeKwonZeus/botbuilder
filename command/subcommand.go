package command

import "github.com/bwmarrin/discordgo"

// Subcommand represents a Discord subcommand.
type Subcommand struct {
	Name           string
	Description    string
	Aliases        []string
	ExecuteCommand bool
	Execute        func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

// SimpleSubcommand creates a Subcommand with the name and the execute function.
func SimpleSubcommand(name string, execute func(*discordgo.Session, *discordgo.MessageCreate, []string)) Subcommand {
	return Subcommand{
		Name:    name,
		Execute: execute,
	}
}
