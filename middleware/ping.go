package middleware

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type pingMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewPingMiddleware(s *discordgo.Session, next MessageMiddleware) MessageMiddleware {
	return pingMiddleware{
		s:    s,
		next: next,
	}
}

func (p pingMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "ping") {
		p.s.ChannelMessageSend(m.ChannelID, "Pong!")
		return nil
	}

	return p.next.Handle(m)
}
