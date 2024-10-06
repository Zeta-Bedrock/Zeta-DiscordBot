package event

import (
	"discord-bot/command"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

func OnInteractionCreate(client *gateway.Session, CmdManager *command.Manager) func(interaction *discord.Interaction) {
	return func(interaction *discord.Interaction) {
		CmdManager.Handler(client)(interaction)
	}
}
