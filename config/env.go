package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var brokerURL string
var username string
var password string
var topic string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	brokerURL = os.Getenv("BROKERURL")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	topic = os.Getenv("TOPIC")
}

func GetBrokerURL() string {
	return brokerURL
}

func GetUsername() string {
	return username
}

func GetPassword() string {
	return password
}

func GetTopic() string {
	return topic
}
