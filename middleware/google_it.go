package middleware

import (
	"strings"
	"net/url"
	"github.com/bwmarrin/discordgo"
)

type googleItMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewGoogleItMiddleware(s *discordgo.Session, next MessageMiddleware) MessageMiddleware {
	return googleItMiddleware{
		s:    s,
		next: next,
	}
}

func (p googleItMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "++google-it") {
		if strings.TrimSpace(m.Content) == "++google-it" {
			p.s.ChannelMessageSend(m.ChannelID, "You should be able to figure it out from here. <https://google.com>")
			return nil
		}
		p.s.ChannelMessageSend(m.ChannelID,
			"Here's a super helpful link! <https://lmgtfy.com/?q="+
				url.QueryEscape(strings.TrimSpace(strings.TrimPrefix(m.Content, "++google-it"))) +
					">")
		return nil
	}

	return p.next.Handle(m)
}
