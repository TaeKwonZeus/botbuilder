# botbuilder

[![Go](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/go.yml/badge.svg)](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/go.yml)

An extension for [discordgo](https://github.com/bwmarrin/discordgo) to create a Discord bot quickly using the Builder pattern.

## Example usage:
```go
package main

import (
	"log"

	"github.com/taekwonzeus/botbuilder"
)

func main() {
	dg, err := botbuilder.NewBotBuilder("token").Build()
	if err != nil {
		log.Fatal("Error when creating bot:", err)
	}

	dg.Open()
}
```
