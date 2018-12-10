package main

import (
	"BotLion/middleware"
	"flag"
	"fmt"
	"os"
	"os/signal"
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

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", *token))
	if err != nil {
		log.Fatal("Error creationg Discord Session", err)
	}

	dg.AddHandler(waitForMessage(dg))

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection ", err)
	}
	fmt.Println("BotLion is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func waitForMessage(s *discordgo.Session) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	finish := middleware.NewFinishMiddleware()
	pongMiddleware := middleware.NewPongMiddleware(s, finish)
	pingMiddleware := middleware.NewPingMiddleware(s, pongMiddleware)
	githubMiddleware := middleware.NewGithubMiddleware(s, pingMiddleware, "https://github.com/camotts/BotLion")
	ignoreBotMiddleware := middleware.NewIgnoreIfFromBotMiddleware(s, githubMiddleware)

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		ignoreBotMiddleware.Handle(m)
	}
}
