package botbuilder

import "github.com/bwmarrin/discordgo"

type CommandHandler interface {
	OnMessageCreate(*discordgo.Session, *discordgo.MessageCreate)
}

type commandHandler struct {
	commands []func(*discordgo.MessageCreate, []string)
}

func (ch *commandHandler) OnMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {

}
