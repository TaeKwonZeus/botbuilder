package command

import "github.com/bwmarrin/discordgo"

// Command represents a Discord command.
type Command struct {
	Name        string
	Description string
	Aliases     []string
	Subcommands []Subcommand
	Execute     func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

// SimpleCommand creates a Command with the name and the execute function.
func SimpleCommand(name string, execute func(*discordgo.Session, *discordgo.MessageCreate, []string)) Command {
	return Command{
		Name:    name,
		Execute: execute,
	}
}
