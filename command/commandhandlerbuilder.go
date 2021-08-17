package command

import "github.com/bwmarrin/discordgo"

type CommandHandlerBuilder interface {
	AddCommand(func(*discordgo.MessageCreate, []string)) CommandHandlerBuilder
	AddCommands(...func(*discordgo.MessageCreate, []string)) CommandHandlerBuilder
	Build() CommandHandler
}

type commandHandlerBuilder struct {
	commands []func(*discordgo.MessageCreate, []string)
}

func (cb *commandHandlerBuilder) AddCommand(command func(*discordgo.MessageCreate, []string)) CommandHandlerBuilder {
	cb.commands = append(cb.commands, command)
	return cb
}

func (cb *commandHandlerBuilder) AddCommands(commands ...func(*discordgo.MessageCreate, []string)) CommandHandlerBuilder {
	cb.commands = append(cb.commands, commands...)
	return cb
}

func (cb *commandHandlerBuilder) Build() CommandHandler {
	return &commandHandler{
		commands: cb.commands,
	}
}
