package command

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

type Manager struct {
	client   *gateway.Session
	commands map[string]Command
}

func NewCommandManager(client *gateway.Session) *Manager {
	return &Manager{
		client:   client,
		commands: make(map[string]Command),
	}
}

func (mgr *Manager) Init() {
	mgr.Register(new(PingCommand))
}

func (mgr *Manager) Handler(client *gateway.Session) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Type != discord.InteractionTypeApplicationCommand {
			return
		}

		if interaction.Member == nil {
			return
		}

		if interaction.Member.User.Bot {
			return
		}

		cmd := mgr.Get(interaction.ApplicationCommandData().Name)

		if cmd != nil {
			_ = client.Interaction.DeferResponse(interaction.Id, interaction.Token, true)

			_ = cmd.Execute(&Context{Client: client, Interaction: interaction, CmdManager: mgr})
		}
	}
}

func (mgr *Manager) Get(name string) Command {
	if cmd, ok := mgr.commands[name]; ok {
		return cmd
	}

	return nil
}

func (mgr *Manager) Register(cmd Command) {
	appCmd := &discord.ApplicationCommand{
		Name:        cmd.Name(),
		Type:        discord.ApplicationCommandChat,
		Description: cmd.Description(),
		Options:     cmd.Options(),
	}

	if _, err := mgr.client.Application.RegisterCommand(mgr.client.Me().Id, "", appCmd); err != nil {
		fmt.Println(err)
	}

	mgr.commands[cmd.Name()] = cmd
}
