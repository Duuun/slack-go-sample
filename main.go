package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}

	token := os.Getenv("SLACK_API_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	api := slack.New(token)
	_, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText("Botによる通知です", false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
