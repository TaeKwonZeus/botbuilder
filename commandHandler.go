package botbuilder

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type commandHandler struct {
	commands []Command
	prefix   string
}

func (ch commandHandler) build(session *discordgo.Session) {
	if ch.commands != nil {
		session.AddHandler(ch.onMessageCreate)
	}
}

func (ch commandHandler) onMessageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.Bot {
		return
	}

	content := messageCreate.Content

	if len(content) <= len(ch.prefix) {
		return
	}

	if content[:len(ch.prefix)] != ch.prefix {
		return
	}

	content = content[len(ch.prefix):]

	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])

	command, err := ch.getCommand(name)
	if err != nil {
		log.Print(err)
		return
	}

	go command.Execute(session, messageCreate, args[1:])
}

func (ch commandHandler) getCommand(name string) (Command, error) {
	for _, command := range ch.commands {
		if command.Name == name {
			return command, nil
		}

		for _, alias := range command.Aliases {
			if alias == name {
				return command, nil
			}
		}
	}

	return Command{}, errors.New("Command not found")
}
