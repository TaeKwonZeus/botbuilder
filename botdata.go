package botbuilder

import "github.com/bwmarrin/discordgo"

// BotData represents Discord bot data and provides another way of using botbuilder.
// Default intents are IntentsGuildMessages.
type BotData struct {
	Token            string
	Prefix           string
	Commands         []Command
	Intents          discordgo.Intent
	EventHandlers    []func(*discordgo.Session, interface{})
	MessageCollector []*discordgo.MessageCreate
}

func (b BotData) Build() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		return nil, err
	}

	for _, i := range b.EventHandlers {
		session.AddHandler(i)
	}

	commandHandler := commandHandler{b.Commands, b.Prefix}
	commandHandler.build(session)

	collectorHandler := collectorHandler{b.MessageCollector}
	collectorHandler.build(session)

	if b.Intents == 0 {
		b.Intents = discordgo.IntentsGuildMessages
	}

	session.Identify.Intents = b.Intents

	return session, nil
}
