# botbuilder

[![](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/go.yml/badge.svg)](https://github.com/TaeKwonZeus/botbuilder/actions/workflows/go.yml)
![](https://img.shields.io/github/issues/TaeKwonZeus/botbuilder)
![](https://img.shields.io/github/forks/TaeKwonZeus/botbuilder)
![](https://img.shields.io/github/stars/TaeKwonZeus/botbuilder)
![](https://img.shields.io/github/license/TaeKwonZeus/botbuilder)
![](https://img.shields.io/github/go-mod/go-version/TaeKwonZeus/botbuilder)

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
