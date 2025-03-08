package config

import (
	"log"
	"msantosfelipe/notification-receiver-lambda/domain"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var ENV domain.Config
var PUSH_NOTIFICATION_ENV domain.PushNotification

func InitVars() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ENV = domain.Config{
		API_PREFIX:     os.Getenv("API_PREFIX"),
		PORT:           os.Getenv("PORT"),
		VALID_API_KEY:  os.Getenv("VALID_API_KEY"),
		ALLOW_ALL_APPS: parseBool(os.Getenv("ALLOW_ANY_APP")),
		ALLOWED_APPS:   parseList(os.Getenv("ALLOWED_APPS")),
		ALLOWED_TITLES: parseList(os.Getenv("ALLOWED_TITLES")),
	}

	PUSH_NOTIFICATION_ENV = domain.PushNotification{
		ENABLE_PUSHOVER:         parseBool(os.Getenv("ENABLE_PUSHOVER_NOTIFICATION")),
		PUSH_OVER_APP_TOKEN:     os.Getenv("PUSH_OVER_APP_TOKEN"),
		PUSH_OVER_APP_RECIPIENT: os.Getenv("PUSH_OVER_APP_RECIPIENT"),
	}
}

func parseList(value string) []string {
	return strings.Split(value, ",")
}

func parseBool(value string) bool {
	return value == "true"
}
