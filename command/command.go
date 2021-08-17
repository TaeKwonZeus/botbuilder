package command

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Subcommands []Subcommand
	Execute     func(*discordgo.MessageCreate, []string)
}

func SimpleCommand(name string, execute func(*discordgo.MessageCreate, []string)) Command {
	return Command{
		Name:    name,
		Execute: execute,
	}
}
