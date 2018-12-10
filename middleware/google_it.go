package middleware

import (
	"strings"

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
		p.s.ChannelMessageSend(m.ChannelID,
			"Here's a super helpful link! <https://lmgtfy.com/?q="+
				QueryEscape(strings.TrimPrefix(m.Content, "++google-it ")+
					">"))
		return nil
	}

	return p.next.Handle(m)
}
