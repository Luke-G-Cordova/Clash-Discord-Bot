package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Luke-G-Cordova/Clash-Discord-Bot/internal/coc"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file if not in production
	if os.Getenv("DEV_MODE") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	clan, err := coc.GetClans(os.Getenv("CLAN_TAG"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(clan)

}

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal("ERROR: creating discord session failed", err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents |= discordgo.IntentsGuildMessages
	discord.Identify.Intents |= discordgo.IntentsGuildMembers

	err = discord.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		log.Println("received a ping")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		log.Println("received a pong")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
