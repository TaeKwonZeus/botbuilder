# botbuilder

[![CI](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/ci.yml/badge.svg)](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/ci.yml)
![issues](https://img.shields.io/github/issues/TaeKwonZeus/botbuilder?logo=github)
![forks](https://img.shields.io/github/forks/TaeKwonZeus/botbuilder?logo=github)
![stars](https://img.shields.io/github/stars/TaeKwonZeus/botbuilder?logo=github)
![license](https://img.shields.io/github/license/TaeKwonZeus/botbuilder)
![Go](https://img.shields.io/github/go-mod/go-version/TaeKwonZeus/botbuilder)
[![Go Reference](https://pkg.go.dev/badge/github.com/taekwonzeus/botbuilder.svg)](https://pkg.go.dev/github.com/taekwonzeus/botbuilder)
[![Go Report Card](https://goreportcard.com/badge/github.com/TaeKwonZeus/botbuilder)](https://goreportcard.com/report/github.com/TaeKwonZeus/botbuilder)

An extension for [discordgo](https://github.com/bwmarrin/discordgo) to create a Discord bot quickly using the Builder pattern.

## Example usage:
```go
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
		AddCommand(ping).
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
```
