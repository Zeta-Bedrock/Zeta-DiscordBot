package context

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

type Context struct {
	CmdManager  CommandManager
	Client      *gateway.Session
	Interaction *discord.Interaction
}

type CommandManager interface {
	Get(name string) Command
}

type Command interface {
	Name() string
	Description() string
	Category() string
	Options() []*discord.ApplicationCommandOption
	Execute(ctx *Context) bool
}
