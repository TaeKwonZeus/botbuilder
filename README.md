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
