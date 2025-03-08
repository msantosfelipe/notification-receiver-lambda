package domain

type Config struct {
	API_PREFIX     string
	PORT           string
	VALID_API_KEY  string
	ALLOW_ALL_APPS bool
	ALLOWED_APPS   []string
	ALLOWED_TITLES []string
}

type PushNotification struct {
	ENABLE_PUSHOVER         bool
	PUSH_OVER_APP_TOKEN     string
	PUSH_OVER_APP_RECIPIENT string
}
