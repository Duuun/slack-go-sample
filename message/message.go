package message

import (
	"flag"
	"fmt"
	"github.com/slack-go/slack"
	"log"

	"github.com/slack-go-sample/utils"
)

func SendMessage() {
	token, channelID := utils.FetchEnv()

	api := slack.New(token)
	_, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText(parseMessage(), false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func parseMessage() string {
	flag.Parse()
	msg := flag.Arg(0)
	return msg
}
