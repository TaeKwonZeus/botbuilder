package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/taekwonzeus/botbuilder"
)

var (
	// Create a command:
	ping = botbuilder.Command{
		Name:        "ping",
		Description: "Ping!",
		Aliases:     []string{"p"},
		Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		},
	}

	// You can also set up a slice where all the messages go.
	messageCollector []*discordgo.MessageCreate
)

func main() {
	// Create a Discord bot
	dg, err := botbuilder.NewBotBuilder(os.Getenv("DISCORD_TOKEN")).
		AddCommands(ping).
		SetMessageCollector(messageCollector).
		Build()
	if err != nil {
		log.Fatal("Error when creating bot: ", err)
	}

	// Open the connection
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	// Close the connection with dg.Close()
}
