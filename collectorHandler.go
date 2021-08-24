package botbuilder

import "github.com/bwmarrin/discordgo"

type collectorHandler struct {
	messageCollector chan *discordgo.MessageCreate
}

func (ch *collectorHandler) onMessageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if ch.messageCollector != nil {
		ch.messageCollector <- messageCreate
	}
}

func (ch *collectorHandler) build(session *discordgo.Session) {
	session.AddHandler(ch.onMessageCreate)
}
