package botbuilder

import "github.com/bwmarrin/discordgo"

// Command represents a Discord command.
type Command struct {
	Name        string
	Description string
	Aliases     []string
	Subcommands []Subcommand
	Execute     func(*discordgo.Session, *discordgo.MessageCreate, []string)
}
