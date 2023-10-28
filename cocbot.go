package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Luke-G-Cordova/Clash-Discord-Bot/internal/coc"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var clan coc.Clan
var cocMemberList []coc.Member
var discordMemberList []*discordgo.Member

func init() {
	// load .env file if not in production
	if os.Getenv("DEV_MODE") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// get current clan info
	_clan, err := coc.GetClan(os.Getenv("CLAN_TAG"))
	if err != nil {
		log.Fatal(err)
	}
	clan = *_clan
	cocMemberList = clan.MemberList
}

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal("ERROR: creating discord session failed", err)
	}

	discord.AddHandler(messageCreate)
	discord.AddHandler(memberReact)

	members, err := discord.GuildMembers(os.Getenv("SERVER_ID"), "0", 100)
	if err != nil {
		log.Fatal("error getting members: ", err)
	}

	discord.Identify.Intents |= discordgo.IntentsGuildMessages
	discord.Identify.Intents |= discordgo.IntentsGuildMembers

	err = discord.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()

}

func setupServer(discord *discordgo.Session) (*discordgo.Channel, error) {
	channels, err := discord.GuildChannels(os.Getenv("SERVER_ID"))
	if err != nil {
		return nil, err
	}

	var channelGroups []*discordgo.Channel
	for _, ch := range channels {
		if ch.Type == discordgo.ChannelTypeGuildCategory {
			channelGroups = append(channelGroups, ch)
		}
	}

	var topGroup *discordgo.Channel
	for _, ch := range channelGroups {
		if strings.ToLower(ch.Name) == "text channels" {
			topGroup = ch
			break
		}
	}
	if topGroup == nil {
		topGroup = channelGroups[0]
	}

	welcome, err := discord.GuildChannelCreate(os.Getenv("SERVER_ID"), "Welcome", discordgo.ChannelTypeGuildText)
	if err != nil {
		return nil, err
	}

	discord.ChannelMessageSend(welcome.ID, "React to one of the names below to change your nick name in this server: ")
	for _, mem := range cocMemberList {
		discord.ChannelMessageSend(welcome.ID, mem.Name)
	}
	return welcome, nil
}

func memberReact(discord *discordgo.Session, m *discordgo.MessageReaction) {
	if m.ChannelID == os.Getenv("WELCOME_CHANNEL_ID") {

	}
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
