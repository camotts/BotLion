package middleware

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type GithubMiddleware struct {
	s         *discordgo.Session
	next      MessageMiddleware
	githubURL string
}

func NewGithubMiddleware(s *discordgo.Session, next MessageMiddleware, githubURL string) MessageMiddleware {
	return GithubMiddleware{s: s, next: next, githubURL: githubURL}
}

func (g GithubMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	if strings.HasPrefix(m.Content, "++github") {
		g.s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You can find me at %s", g.githubURL))
		return nil
	}
	return g.next.Handle(m)
}
