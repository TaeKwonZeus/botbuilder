package botbuilder

import "github.com/bwmarrin/discordgo"

type slashCommandHandler struct {
	slashCommands        []*discordgo.ApplicationCommand
	slashCommandHandlers map[string]SlashCommandFunction
}

func (sch slashCommandHandler) build(session *discordgo.Session) {
	if sch.slashCommands != nil {
		session.AddHandler(sch.onInteractionCreate)
	}
}

func (sch slashCommandHandler) onInteractionCreate(session *discordgo.Session, event *discordgo.InteractionCreate) {

}
