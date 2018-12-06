package middleware

import "github.com/bwmarrin/discordgo"

type ignoreIfFromBotMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewIgnoreIfFromBotMiddleware(s *discordgo.Session, next MessageMiddleware) MessageMiddleware {
	return ignoreIfFromBotMiddleware{
		s:    s,
		next: next,
	}
}

func (p ignoreIfFromBotMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if m.Author.ID != p.s.State.User.ID {
		p.next.Handle(m)
	}
	return nil
}
