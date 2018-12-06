package middleware

import "github.com/bwmarrin/discordgo"

type MessageMiddleware interface {
	Handle(*discordgo.MessageCreate) MessageMiddleware
}
