package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/taekwonzeus/botbuilder"
)

func main() {
	// Create a command
	ping := botbuilder.Command{
		Name:        "ping",
		Description: "Ping!",
		Aliases:     []string{"p"},
		Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		},
	}

	// You can also set up a slice where all the messages go
	var messageCollector []*discordgo.MessageCreate

	// Create a Discord bot
	dg, err := botbuilder.BotData{
		Token:            os.Getenv("DISCORD_TOKEN"),
		Prefix:           "!",
		Commands:         []botbuilder.Command{ping},
		MessageCollector: messageCollector,
	}.Build()
	if err != nil {
		log.Println(err)
		return
	}

	// Open the connection
	err = dg.Open()
	if err != nil {
		log.Println(err)
		return
	}

	// When you exit, call dg.Close()
}
