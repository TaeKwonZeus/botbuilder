package botbuilder

import "github.com/bwmarrin/discordgo"

func NewBotBuilder(token string) BotBuilder {
	return &botBuilder{token: token}
}

type BotBuilder interface {
	AddEventHandler(func(*discordgo.Session, interface{})) BotBuilder
	AddEventHandlers(...func(*discordgo.Session, interface{})) BotBuilder
	AddCommandHandler(CommandHandler) BotBuilder
	Build() (*discordgo.Session, error)
}

type botBuilder struct {
	token          string
	eventHandlers  []func(*discordgo.Session, interface{})
	commandHandler CommandHandler
}

func (bb *botBuilder) AddEventHandler(eventHandler func(*discordgo.Session, interface{})) BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandler)
	return bb
}

func (bb *botBuilder) AddEventHandlers(eventHandlers ...func(*discordgo.Session, interface{})) BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandlers...)
	return bb
}

func (bb *botBuilder) AddCommandHandler(commandHandler CommandHandler) BotBuilder {
	bb.commandHandler = commandHandler
	return bb
}

func (bb *botBuilder) Build() (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + bb.token)
	if err != nil {
		return nil, err
	}

	for _, i := range bb.eventHandlers {
		dg.AddHandler(i)
	}

	if bb.commandHandler != nil {
		dg.AddHandler(bb.commandHandler.OnMessageCreate)
	}

	return dg, nil
}
