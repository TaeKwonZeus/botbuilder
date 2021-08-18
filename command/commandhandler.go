// Package command adds command handlers functionality.
package command

import "github.com/bwmarrin/discordgo"

// NewCommandHandler creates a CommandHandler with the commands specified.
func NewCommandHandler(commands ...Command) CommandHandler {
	return &commandHandler{commands: commands}
}

// CommandHandler represents a Discord command handler.
// You can create your own struct based on CommandHandler and set it as the command handler.
type CommandHandler interface {
	OnMessageCreate(*discordgo.Session, *discordgo.MessageCreate)
}

type commandHandler struct {
	commands []Command
}

// OnMessageCreate is executed by the command handler when a message is created.
func (ch *commandHandler) OnMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
}
