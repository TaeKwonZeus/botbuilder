package botbuilder

import "github.com/bwmarrin/discordgo"

type slashCommandHandler struct {
	slashCommands        []*discordgo.ApplicationCommand
	slashCommandHandlers map[string]SlashCommandFunction
}
