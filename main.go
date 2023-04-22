package main

import (
	"os/signal"
	"syscall"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/aricodes-oss/std"
	"github.com/nicklaw5/helix/v2"

	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"

	"whisprite/core"
	"whisprite/handlers"
)

var log = std.Logger

var (
	CLIENT_ID    = os.Getenv("TWITCH_CLIENT_ID")
	ACCESS_TOKEN = os.Getenv("TWITCH_ACCESS_TOKEN")
	PASSWORD     = os.Getenv("TWITCH_PASSWORD")
	USERNAME     = os.Getenv("TWITCH_USERNAME")
	CHANNELS     = strings.Split(os.Getenv("TWITCH_CHANNEL_NAMES"), ",")
	BOT_MODE     = os.Getenv("BOT_MODE")
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	if BOT_MODE == "development" {
		log.WithDebug()
	}

	log.Info("Connecting to twitch services...")
	helixClient, err := helix.NewClient(&helix.Options{ClientID: CLIENT_ID, UserAccessToken: ACCESS_TOKEN})
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Bootstrapping the application...")
	dispatch := &core.Dispatch{Twitch: helixClient}
	dispatch.Register(handlers.Counters...)

	log.Info("Opening up the writer thread...")
	writer := &irc.Conn{}
	writer.SetLogin(USERNAME, PASSWORD)
	if err := writer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	/*
	 * This is assigned to its own variable to avoid getting nil dereference errors
	 * when referring to the writer inside of the message handler. Anonymous functions
	 * declared as variables are closures that evaluate their scope at run time, whereas
	 * anonymous functions passed directly as arguments evaluate their scope at compile time
	 */
	onMsg := func(shardID int, msg irc.ChatMessage) {
		log.Debugf("[%v] %v", shardID, msg)

		isCommand := msg.Text[0] == '!'
		if isCommand {
			event := &core.Event{
				ChatMessage:   msg,
				IsMod:         msg.Sender.IsModerator,
				IsVIP:         msg.Sender.IsVIP,
				IsBroadcaster: msg.Sender.IsBroadcaster,

				Wsay:  writer.Say,
				Wsayf: writer.Sayf,
			}

			err := event.Parse()
			if err != nil {
				panic(err)
			}

			go dispatch.Handle(event)
		}
	}

	log.Info("Creating sharded read pool...")
	reader := twitch.IRC()
	reader.OnShardMessage(onMsg)
	defer reader.Close()

	log.Info("Pulling user information from the API...")
	resp, err := helixClient.GetUsers(&helix.UsersParams{Logins: CHANNELS})
	if err != nil {
		log.Info(resp)
		log.Fatal(err)
	}

	for _, user := range resp.Data.Users {
		log.Infof("Joining room %s (%s)...", user.DisplayName, user.ID)
		reader.Join(user.Login)
	}

	log.Info("Ret-2-Go!")

	<-sc
	log.Info("Shutting down! Have a good night :)")
}
