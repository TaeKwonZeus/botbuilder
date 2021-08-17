package command

import "github.com/bwmarrin/discordgo"

type Subcommand struct {
	Name           string
	Description    string
	Aliases        []string
	ExecuteCommand bool
	Execute        func(*discordgo.MessageCreate, []string)
}

func SimpleSubcommand(name string, execute func(*discordgo.MessageCreate, []string)) Subcommand {
	return Subcommand{}
}
