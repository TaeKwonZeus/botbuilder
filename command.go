package botbuilder

import "github.com/bwmarrin/discordgo"

// CommandFunction represents a function executed on a command call.
type CommandFunction func(session *discordgo.Session, event *discordgo.MessageCreate, args []string)

// Command represents a Discord command.
type Command struct {
	Name        string
	Description string
	Aliases     []string
	Subcommands []Subcommand
	Execute     CommandFunction
}
