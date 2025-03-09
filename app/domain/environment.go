package domain

type Config struct {
	VALID_API_KEY  string
	ALLOW_ALL_APPS bool
	ALLOWED_APPS   []string
	ALLOWED_TITLES []string
}

type PushNotification struct {
	PUSH_OVER_APP_TOKEN     string
	PUSH_OVER_APP_RECIPIENT string
}
