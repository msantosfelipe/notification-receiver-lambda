package config

import (
	"encoding/json"
	"fmt"
	"msantosfelipe/notification-receiver-lambda/internal/domain"
	"os"

	"github.com/joho/godotenv"
)

var ENV domain.Config
var PushNotificationEnv domain.PushNotification

func InitVars() {
	// Load .env file
	// local will use .env, in AWS will use os env vars
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	ENV = domain.Config{
		IsLocal:      parseBool(os.Getenv("IS_LOCAL")),
		ValidApiKey:  os.Getenv("VALID_API_KEY"),
		AllowAllApps: parseBool(os.Getenv("ALLOW_ANY_APP")),
		AppsAllowed:  parseApps(os.Getenv("APPS_ALLOWED_JSON")),
	}

	PushNotificationEnv = domain.PushNotification{
		PushOverAppToken:     os.Getenv("PUSH_OVER_APP_TOKEN"),
		PushOverAppRecipient: os.Getenv("PUSH_OVER_APP_RECIPIENT"),
	}
}

func parseApps(appsAllowedJson string) []domain.AppAllowed {
	var appsAllowed []domain.AppAllowed

	err := json.Unmarshal([]byte(appsAllowedJson), &appsAllowed)
	if err != nil {
		fmt.Println("Error parsing APPS_ALLOWED_JSON:", err)
	}
	return appsAllowed
}

func parseBool(value string) bool {
	return value == "true"
}
