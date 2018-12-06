package core

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageMiddleware interface {
	Handle(*discordgo.MessageCreate) MessageMiddleware
}

type IgnoreIfFromBotMiddleware struct {
	S    *discordgo.Session
	Next MessageMiddleware
}

func (p IgnoreIfFromBotMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if m.Author.ID != p.S.State.User.ID {
		p.Next.Handle(m)
	}
	return nil
}

type PingMiddleWare struct {
	S    *discordgo.Session
	Next MessageMiddleware
}

func (p PingMiddleWare) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "ping") {
		p.S.ChannelMessageSend(m.ChannelID, "Pong!")
		return nil
	}

	return p.Next.Handle(m)
}

type FinishMiddleware struct {
	S    *discordgo.Session
	Next MessageMiddleware
}

func (p FinishMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	return nil
}
