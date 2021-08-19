# botbuilder

[![CI](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/ci.yml/badge.svg)](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/ci.yml)
![issues](https://img.shields.io/github/issues/TaeKwonZeus/botbuilder?logo=github)
![forks](https://img.shields.io/github/forks/TaeKwonZeus/botbuilder?logo=github)
![stars](https://img.shields.io/github/stars/TaeKwonZeus/botbuilder?logo=github)
![license](https://img.shields.io/github/license/TaeKwonZeus/botbuilder)
![Go](https://img.shields.io/github/go-mod/go-version/TaeKwonZeus/botbuilder)
[![Go Reference](https://pkg.go.dev/badge/github.com/taekwonzeus/botbuilder.svg)](https://pkg.go.dev/github.com/taekwonzeus/botbuilder)

An extension for [discordgo](https://github.com/bwmarrin/discordgo) to create a Discord bot quickly using the Builder pattern.

## Example usage:
```go
package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/taekwonzeus/botbuilder"
	"github.com/taekwonzeus/botbuilder/command"
)

var (
	// Create a fully tuned command:
	ping = command.Command{
		Name:        "ping",
		Description: "Ping!",
		Aliases:     []string{"p"},
		Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		},
	}

	// Or a simple one:
	test = command.SimpleCommand(
		"test",
		func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(m.ChannelID, "test response")
		},
	)
)

func main() {
	// Create a command handler
	handler := command.NewCommandHandler(ping, test)

	// Create a Discord bot
	dg, err := botbuilder.NewBotBuilder("token").SetCommandHandler(handler).Build()
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
