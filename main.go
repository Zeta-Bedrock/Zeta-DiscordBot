package main

import (
	"discord-bot/command"
	botevent "discord-bot/event"
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/Goscord/goscord/goscord/gateway/event"
	"github.com/joho/godotenv"

	"os"
)

var (
	client     *gateway.Session
	cmdManager *command.Manager
)

func main() {
	_ = godotenv.Load()

	client := goscord.New(&gateway.Options{
		Token:   os.Getenv("TOKEN"),
		Intents: gateway.IntentsAll,
	})

	cmdManager = command.NewCommandManager(client)

	_ = client.On(event.EventReady, botevent.OnReady(client, cmdManager))
	_ = client.On(event.EventInteractionCreate, botevent.OnInteractionCreate(client, cmdManager))

	err := client.Login()
	if err != nil {
		return
	}

	select {}
}
