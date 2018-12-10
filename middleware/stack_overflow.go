package middleware

import (
	"strings"
	"net/url"
	"github.com/bwmarrin/discordgo"
)

type stackOverflowMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewStackOverflowMiddleware(s *discordgo.Session, next MessageMiddleware) MessageMiddleware {
	return stackOverflowMiddleware{
		s:    s,
		next: next,
	}
}

func (p stackOverflowMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "++stack-overflow") {
		if strings.TrimSpace(m.Content) == "++stack-overflow" {
			p.s.ChannelMessageSend(m.ChannelID, "Well... I found <https://stackoverflow.com>")
			return nil
		}
		p.s.ChannelMessageSend(m.ChannelID,
			"This topic might help you: <https://stackoverflow.com/search?q="+
				url.QueryEscape(strings.TrimSpace(strings.TrimPrefix(m.Content, "++stack-overflow"))) +
				">")
		return nil
	}

	return p.next.Handle(m)
}
