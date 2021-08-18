// Package botbuilder allows for simple Discord bot initialization utilizing the builder pattern.
// The Discord API functionality is provided by a well-known package discordgo.
package botbuilder

import (
	"github.com/bwmarrin/discordgo"
	"github.com/taekwonzeus/botbuilder/command"
)

// NewBotBuilder creates a new BotBuilder object and writes the token specified into the builder.
// You don't have to add "Bot " to the token like in discordgo as the builder will do it for you.
func NewBotBuilder(token string) BotBuilder {
	return &botBuilder{token: token}
}

// BotBuilder represents a Discord bot builder.
type BotBuilder interface {
	AddEventHandler(func(*discordgo.Session, interface{})) BotBuilder
	AddEventHandlers(...func(*discordgo.Session, interface{})) BotBuilder
	SetCommandHandler(command.CommandHandler) BotBuilder
	SetIntents(discordgo.Intent) BotBuilder
	Build() (*discordgo.Session, error)
}

type botBuilder struct {
	token          string
	eventHandlers  []func(*discordgo.Session, interface{})
	commandHandler command.CommandHandler
	intents        discordgo.Intent
}

// AddEventHandler adds a singular Discord event handler to the bot.
func (bb *botBuilder) AddEventHandler(eventHandler func(*discordgo.Session, interface{})) BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandler)
	return bb
}

// AddEventHandlers adds multiple Discord event handlers to the bot.
func (bb *botBuilder) AddEventHandlers(eventHandlers ...func(*discordgo.Session, interface{})) BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandlers...)
	return bb
}

// SetCommandHandler sets the bot's command handler.
// Additional SetCommander calls will rewrite the CommandHandler.
func (bb *botBuilder) SetCommandHandler(commandHandler command.CommandHandler) BotBuilder {
	bb.commandHandler = commandHandler
	return bb
}

func (bb *botBuilder) SetIntents(intents discordgo.Intent) BotBuilder {
	bb.intents = intents
	return bb
}

// Build creates a Discord bot, adding the command handler and event handlers specified.
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

	if bb.intents == 0 {
		bb.intents = discordgo.IntentsAllWithoutPrivileged
	}

	dg.Identify.Intents = bb.intents

	return dg, nil
}
