package middleware

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type pongMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewPongMiddleware(s *discordgo.Session, next MessageMiddleware) MessageMiddleware {
	return pongMiddleware{
		s:    s,
		next: next,
	}
}

func (p pongMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "pong") {
		p.s.ChannelMessageSend(m.ChannelID, "Ping!")
		return nil
	}

	return p.next.Handle(m)
}
