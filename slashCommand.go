package botbuilder

import "github.com/bwmarrin/discordgo"

type SlashCommand struct {
	ApplicationCommand *discordgo.ApplicationCommand
	Execute            func(session *discordgo.Session, event *discordgo.InteractionCreate)
	Guild              string
}
