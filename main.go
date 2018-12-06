package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	//move to more feature filled logging package
	"log"

	"github.com/bwmarrin/discordgo"
)

func main() {
	//flag parsing
	//be more smart about this later
	token := flag.String("t", "", "Bot Authorization Token")
	flag.Parse()

	dg, err := discordgo.New(fmt.Sprintf("Bot %v", token))
	if err != nil {
		log.Fatal("Error creationg Discord Session", err)
	}

	dg.AddHandlerOnce(waitForMessage)

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection", err)
	}
	fmt.Println("BotLion is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func waitForMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	msg := strings.ToLower(m.Content)
	if strings.HasPrefix(msg, "ping") {
		for i := 0; i < strings.Count(msg, "s")+1; i++ {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}
	}

	if strings.HasPrefix(msg, "pong") {
		for i := 0; i < strings.Count(msg, "s")+1; i++ {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}
	}
}
