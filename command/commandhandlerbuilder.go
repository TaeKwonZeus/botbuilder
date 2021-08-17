package command

func NewCommandHandlerBuilder() CommandHandlerBuilder {
	return &commandHandlerBuilder{}
}

type CommandHandlerBuilder interface {
	AddCommand(Command) CommandHandlerBuilder
	AddCommands(...Command) CommandHandlerBuilder
	Build() CommandHandler
}

type commandHandlerBuilder struct {
	commands []Command
}

func (cb *commandHandlerBuilder) AddCommand(command Command) CommandHandlerBuilder {
	cb.commands = append(cb.commands, command)
	return cb
}

func (cb *commandHandlerBuilder) AddCommands(commands ...Command) CommandHandlerBuilder {
	cb.commands = append(cb.commands, commands...)
	return cb
}

func (cb *commandHandlerBuilder) Build() CommandHandler {
	return &commandHandler{
		commands: cb.commands,
	}
}
