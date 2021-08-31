package botbuilder

import (
	"github.com/bwmarrin/discordgo"
)

type slashCommandHandler struct {
	slashCommands        []*discordgo.ApplicationCommand
	slashCommandHandlers map[string]SlashCommandFunction
	slashCommandGuilds   map[string]string
}

func (sch slashCommandHandler) build(session *discordgo.Session) {
	if sch.slashCommands != nil {
		for _, i := range sch.slashCommands {
			session.ApplicationCommandCreate(session.State.User.ID, sch.slashCommandGuilds[i.Name], i)
		}
		session.AddHandler(sch.onInteractionCreate)
	}
}

func (sch slashCommandHandler) onInteractionCreate(session *discordgo.Session, event *discordgo.InteractionCreate) {
	if h, ok := sch.slashCommandHandlers[event.ApplicationCommandData().Name]; ok {
		h(session, event)
	}
}
