package middleware

import (
	"github.com/bwmarrin/discordgo"
)

func AddPollHandler(s *discordgo.Session) {
	messageCreate := func(s *discordgo.Session, m *discordgo.MessageCreate) {

	}
	s.AddHandler(messageCreate)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
