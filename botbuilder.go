// Package botbuilder allows for simple Discord bot initialization utilizing the builder pattern.
// The Discord API functionality is provided by a well-known package discorsessiono.
package botbuilder

import (
	"github.com/bwmarrin/discordgo"
)

// NewBotBuilder creates a new BotBuilder object and writes the token specified into the builder.
// You don't have to add "Bot " to the token like in discorsessiono as the builder will do it for you.
func NewBotBuilder(token string) *BotBuilder {
	return &BotBuilder{token: token}
}

// BotBuilder represents a Discord bot builder.
type BotBuilder struct {
	token               string
	eventHandlers       []func(session *discordgo.Session, event interface{})
	commandHandler      commandHandler
	slashCommandHandler slashCommandHandler
	intents             discordgo.Intent
	collectorHandler    collectorHandler
}

// AddEventHandler adds a single Discord event handler to the bot.
func (bb *BotBuilder) AddEventHandler(eventHandler func(*discordgo.Session, interface{})) *BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandler)
	return bb
}

// AddEventHandlers adds multiple Discord event handlers to the bot.
func (bb *BotBuilder) AddEventHandlers(eventHandlers ...func(*discordgo.Session, interface{})) *BotBuilder {
	bb.eventHandlers = append(bb.eventHandlers, eventHandlers...)
	return bb
}

// AddCommand adds a single command to the bot.
func (bb *BotBuilder) AddCommand(command *Command) *BotBuilder {
	bb.commandHandler.commands = append(bb.commandHandler.commands, command)
	return bb
}

// AddCommands adds multiple commands to the bot.
func (bb *BotBuilder) AddCommands(commands ...*Command) *BotBuilder {
	bb.commandHandler.commands = append(bb.commandHandler.commands, commands...)
	return bb
}

// AddSlashCommand adds a single slash command to the bot.
func (bb *BotBuilder) AddSlashCommand(slashCommand *SlashCommand) *BotBuilder {
	bb.slashCommandHandler.slashCommands = append(bb.slashCommandHandler.slashCommands, slashCommand.ApplicationCommand)
	bb.slashCommandHandler.slashCommandHandlers[slashCommand.ApplicationCommand.Name] = slashCommand.Execute
	return bb
}

// AddSlashCommands adds multiple slash commands to the bot.
func (bb *BotBuilder) AddSlashCommands(slashCommands ...*SlashCommand) *BotBuilder {
	for _, i := range slashCommands {
		bb.AddSlashCommand(i)
	}

	return bb
}

// SetIntents sets the bot's intents.
// If not called, the bot's intents will be set to IntentsGuildMessages.
func (bb *BotBuilder) SetIntents(intents discordgo.Intent) *BotBuilder {
	bb.intents = intents
	return bb
}

// SetMessageBuffer sets a message collector for the bot.
// Every time a MessageCreate event is invoked the event data is passed into c.
func (bb *BotBuilder) SetMessageCollector(collector []*discordgo.MessageCreate) *BotBuilder {
	bb.collectorHandler.messageCollector = collector
	return bb
}

// SetPrefix sets a prefix for the bot.
// By default, there is no prefix.
func (bb *BotBuilder) SetPrefix(prefix string) *BotBuilder {
	bb.commandHandler.prefix = prefix
	return bb
}

// Build creates a Discord bot, adding the command handler and event handlers specified.
func (bb *BotBuilder) Build() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + bb.token)
	if err != nil {
		return nil, err
	}

	for _, i := range bb.eventHandlers {
		if i != nil {
			session.AddHandler(i)
		}
	}

	bb.collectorHandler.build(session)
	bb.commandHandler.build(session)
	bb.slashCommandHandler.build(session)

	if bb.intents == 0 {
		bb.intents = discordgo.IntentsGuildMessages
	}

	session.Identify.Intents = bb.intents

	return session, nil
}
