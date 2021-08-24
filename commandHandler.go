package botbuilder

import "github.com/bwmarrin/discordgo"

type commandHandler struct {
	commands []*Command
}

func (ch *commandHandler) build(session *discordgo.Session) {
	if ch.commands != nil {
		session.AddHandler(ch.onMessageCreate)
	}
}

func (ch *commandHandler) onMessageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
}
