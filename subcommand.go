package botbuilder

import "github.com/bwmarrin/discordgo"

// Subcommand represents a Discord subcommand.
type Subcommand struct {
	Name           string
	Description    string
	Aliases        []string
	ExecuteCommand bool
	Execute        func(*discordgo.Session, *discordgo.MessageCreate, []string)
}
