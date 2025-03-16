package config

import (
	"encoding/json"
	"fmt"
	"msantosfelipe/notification-receiver-lambda/domain"
	"os"

	"github.com/joho/godotenv"
)

var ENV domain.Config
var PUSH_NOTIFICATION_ENV domain.PushNotification

func InitVars() {
	// Load .env file
	// local will use .env, in AWS will use os env vars
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	ENV = domain.Config{
		IS_LOCAL:       parseBool(os.Getenv("IS_LOCAL")),
		VALID_API_KEY:  os.Getenv("VALID_API_KEY"),
		ALLOW_ALL_APPS: parseBool(os.Getenv("ALLOW_ANY_APP")),
		APPS_ALLOWED:   parseApps(os.Getenv("APPS_ALLOWED_JSON")),
	}

	PUSH_NOTIFICATION_ENV = domain.PushNotification{
		PUSH_OVER_APP_TOKEN:     os.Getenv("PUSH_OVER_APP_TOKEN"),
		PUSH_OVER_APP_RECIPIENT: os.Getenv("PUSH_OVER_APP_RECIPIENT"),
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
