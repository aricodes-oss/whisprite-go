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

func commandDispatch(c chan *CommandEvent, running chan bool) {
	for {
		select {
		case ev := <-c:
			log.Debugf("Received event %v", ev)
			err := ev.Parse()

			if err != nil {
				log.Errorf("Failed to parse message %s | %v", ev.Text, err)
			}
		case <-running:
			return
		}
	}
}

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	dispatch := make(chan *CommandEvent, 1024)
	running := make(chan bool, 1)

	go commandDispatch(dispatch, running)
	defer close(running)

	if BOT_MODE == "development" {
		log.WithDebug()
	}

	log.Info("Booting whisprite! Connecting to twitch services...")
	helixClient, err := helix.NewClient(&helix.Options{ClientID: CLIENT_ID, UserAccessToken: ACCESS_TOKEN})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Done! Bootstrapping the application")

	log.Info("Creating IRC read pool...")
	reader := twitch.IRC()
	reader.OnShardMessage(func(shardID int, msg irc.ChatMessage) {
		log.Debugf("[%v] %v", shardID, msg)
		isCommand := msg.Text[0] == '!'

		if isCommand {
			dispatch <- &CommandEvent{
				msg,
				"",
				[]string{},
				helixClient,
			}
		}
	})
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

	log.Info("Opening up the writer thread...")
	writer := &irc.Conn{}
	writer.SetLogin(USERNAME, PASSWORD)
	if err := writer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	log.Info("Ret-2-Go!")

	<-sc
	log.Info("Shutting down! Have a good night :)")
}
