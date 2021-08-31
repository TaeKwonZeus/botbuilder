package botbuilder

import "github.com/bwmarrin/discordgo"

type SlashCommandFunction func(session *discordgo.Session, event *discordgo.InteractionCreate)

type SlashCommand struct {
	ApplicationCommand *discordgo.ApplicationCommand
	Execute            SlashCommandFunction
}
