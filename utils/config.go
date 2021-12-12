package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FetchEnv() (string, string) {
	initEnv()
	token := os.Getenv("SLACK_API_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	return token, channelID
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
