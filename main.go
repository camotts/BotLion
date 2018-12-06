package main

import (
	"BotLion/core"
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

	dg.AddHandler(waitForMessage)

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

func waitForMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	go handleMessage(s, m)
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	finish := core.FinishMiddleware{
		S:    s,
		Next: nil,
	}

	pingMiddleware := core.PingMiddleWare{
		S:    s,
		Next: finish,
	}

	ignoreBotMiddleware := core.IgnoreIfFromBotMiddleware{
		S:    s,
		Next: pingMiddleware,
	}

	ignoreBotMiddleware.Handle(m)
}
