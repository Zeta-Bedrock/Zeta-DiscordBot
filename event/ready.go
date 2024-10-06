package event

import (
	"discord-bot/command"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/JustTimmm/GoColor"
)

func OnReady(client *gateway.Session, CmdManager *command.Manager) func() {
	return func() {
		GoColor.ColorLog(GoColor.ColorOption{TextColor: GoColor.Green}, "Bot started 🚀\n")
		GoColor.WarnLog("Logged in as %s ✨ \n", client.Me().Username)

		CmdManager.Init()

		_ = client.SetActivity(&discord.Activity{Name: "/help", Type: discord.ActivityListening})
	}
}
