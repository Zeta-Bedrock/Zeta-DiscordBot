package moderation

import (
	"discord-bot/command/context"
	"discord-bot/utils"
	"github.com/Goscord/goscord/goscord/discord"
)

type KickCommand struct{}

func (c *KickCommand) Name() string {
	return "kick"
}

func (c *KickCommand) Description() string {
	return "Kick user"
}

func (c *KickCommand) Category() string {
	return "moderation"
}

func (c *KickCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "user",
			Type:        discord.ApplicationCommandOptionUser,
			Description: "User",
			Required:    true,
		},
	}
}

func (c *KickCommand) Execute(ctx *context.Context) bool {
	if !ctx.Interaction.Member.Permissions.Has(utils.PermKick) {
		_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, "You don't have the permissions 🤡")
	} else {

	}

	return true
}
