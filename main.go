package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/tjtjtjtj/go-slack/slack"
)

const (
	slackAPIURL = "https://slack.com/api"
)

type envConfig struct {

	// GHEToken is bot user token to access to GHE API.
	GHEToken string `envconfig:"GHE_TOKEN" required:"false"`
}

var env envConfig

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		return 1
	}

	log.Printf("env:%v", env)

	// Listening slack event and response
	log.Printf("[INFO] Start slack event listening")
	client := slack.New(env.BotToken)
	slackListener := &SlackListener{
		client:    client,
		botID:     env.BotID,
		channelID: env.ChannelID,
	}

	return 0
}
