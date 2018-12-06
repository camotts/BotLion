package middleware

import "github.com/bwmarrin/discordgo"

type finishMiddleware struct {
	s    *discordgo.Session
	next MessageMiddleware
}

func NewFinishMiddleware() MessageMiddleware {
	return finishMiddleware{}
}

func (p finishMiddleware) Handle(m *discordgo.MessageCreate) MessageMiddleware {
	return nil
}
